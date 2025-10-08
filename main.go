package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
)

type CurrencyRate struct {
	BankName string
	BuyRate  string
	SellRate string
	Date     string
}

func main() {
	var rates []CurrencyRate

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Could not start Playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("Could not launch Firefox: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("Could not create page: %v", err)
	}

	if _, err = page.Goto("https://www.banki.ru/products/currency/?source=main_exchange_rates_converter"); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}

	page.WaitForTimeout(5000)

	// Ищем все контейнеры с курсами валют
	currencyContainers := page.Locator("//div[contains(@class, 'resultItemstyled__StyledWrapperResult')]")
	containerCount, _ := currencyContainers.Count()

	fmt.Printf("Найдено контейнеров с курсами: %d\n", containerCount)

	for i := 0; i < containerCount; i++ {
		container := currencyContainers.Nth(i)

		// Название банка
		bankNameLocator := container.Locator("[data-test='currenct--result-item--name']")
		bankName, err := bankNameLocator.TextContent()
		if err != nil || bankName == "" {
			continue
		}
		bankName = strings.TrimSpace(bankName)

		// Ищем все элементы с data-test='text' в этом контейнере
		allTextElements := container.Locator("[data-test='text']")
		elementCount, _ := allTextElements.Count()

		var buyRate, sellRate string

		// Проходим по всем элементам и определяем их назначение по тексту
		for j := 0; j < elementCount; j++ {
			text, err := allTextElements.Nth(j).TextContent()
			if err != nil {
				continue
			}
			text = strings.TrimSpace(text)

			// Если это надпись "Покупка", то следующий элемент - курс покупки
			if text == "Покупка" && j+1 < elementCount {
				buyRate, _ = allTextElements.Nth(j + 1).TextContent()
				buyRate = strings.TrimSpace(buyRate)
			}

			// Если это надпись "Продажа", то следующий элемент - курс продажи
			if text == "Продажа" && j+1 < elementCount {
				sellRate, _ = allTextElements.Nth(j + 1).TextContent()
				sellRate = strings.TrimSpace(sellRate)
			}
		}

		rate := CurrencyRate{
			BankName: bankName,
			BuyRate:  buyRate,
			SellRate: sellRate,
			Date:     time.Now().Format("2006-01-02"),
		}

		rates = append(rates, rate)
	}

	// Вывод результатов
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("КУРСЫ ВАЛЮТ НА", time.Now().Format("02.01.2006"))
	fmt.Println(strings.Repeat("=", 50))

	for i, rate := range rates {
		fmt.Printf("%2d. %-30s Покупка: %-10s Продажа: %-10s\n",
			i+1, rate.BankName, rate.BuyRate, rate.SellRate)
	}

	fmt.Printf("\nВсего обработано: %d банков\n", len(rates))
}

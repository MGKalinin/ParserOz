package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	const (
		seleniumPath = "/opt/homebrew/bin/chromedriver"
		port         = 4444
	)

	// Запускаем ChromeDriver сервер
	service, err := selenium.NewChromeDriverService(seleniumPath, port)
	if err != nil {
		log.Fatal("❌ Ошибка запуска ChromeDriver:", err)
	}
	defer service.Stop()

	fmt.Println("✅ ChromeDriver сервер запущен")

	// Простые настройки без сложных опций
	caps := selenium.Capabilities{"browserName": "chrome"}

	// Подключаемся к WebDriver
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatal("❌ Ошибка подключения к WebDriver:", err)
	}
	defer wd.Quit()

	fmt.Println("✅ Браузер запущен")

	// Переходим на страницу
	url := "https://www.ozon.ru/product/muzhskie-naruchnye-chasy-casio-g-shock-gw-9400-1er-2058384940/"
	fmt.Printf("🌐 Загружаем страницу: %s\n", url)

	err = wd.Get(url)
	if err != nil {
		log.Fatal("❌ Ошибка загрузки страницы:", err)
	}

	// Ждем загрузки
	fmt.Println("⏳ Ожидаем загрузки страницы...")
	time.Sleep(5 * time.Second)

	// Получаем цену с надежным селектором
	fmt.Println("🔍 Ищем цену...")
	price, err := wd.FindElement(selenium.ByCSSSelector, "span[data-widget='webPrice']")
	if err != nil {
		log.Fatal("❌ Не удалось найти элемент цены:", err)
	}

	// Получаем текст цены
	priceText, err := price.Text()
	if err != nil {
		log.Fatal("❌ Не удалось получить текст цены:", err)
	}

	fmt.Printf("\n🎯 ЦЕНА НАЙДЕНА: %s\n", priceText)
	fmt.Printf("📎 URL товара: %s\n", url)
}

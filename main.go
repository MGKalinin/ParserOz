package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// здесь вызвать парсер по такому-то адресу
// указать путь сохранения в файл
func main() {
	fName := "ozonePrice.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//указать поля
	writer.Write([]string{"Name", "Price", "Date"})
	c := colly.NewCollector()
	c.OnHTML("#currencies-all tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".currency-name-container"),
			e.ChildText(".col-symbol"),
			e.ChildAttr("a.price", "data-usd"),
		})
	})
	c.Visit("https://www.ozon.ru/category/aksessuary-7697/?category_was_predicted=true&deny_category_prediction=true&from_global=true&text=casio+gw+9400")
	log.Printf("Scraping finished, check file %q for results\n", fName)

}

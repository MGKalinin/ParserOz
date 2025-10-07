package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Product data structure to keep the scraped data
type Product struct {
	Price string
}

func main() {
	// initialize the slice of structs that will contain the scraped data
	var products []Product
	// instantiate a new collector object
	c := colly.NewCollector()
	// OnHTML callback
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// initialize a new Product instance
		product := Product{}
		// scrape the target data
		product.Price = e.ChildText("/html/body/div[2]/main/div[2]/div[1]/div/div[2]/div/div[3]/div[3]/div/div[1]/div[1]/div/div/div/div/div/span/ins")

		// add the product instance with scraped data to the list of products
		products = append(products, product)

	})
	// open the target URL
	c.Visit("https://www.wildberries.ru/catalog/75455564/detail.aspx")
	// Print results
	for _, product := range products {
		fmt.Printf("Price: %s\n", product.Price)
	}
}

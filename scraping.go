package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fetchURL := "https://www.orami.co.id/shopping/promo/store-405640-terbaru"
	fileName := "orami.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("ERROR: Could not create file : \n", fileName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write column headers of the text file
	writer.Write([]string{"Nama", "Gambar", "Harga"})

	// Instantiate the default Collector
	c := colly.NewCollector()

	// Before making a request, print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnHTML(`.px-8 pb-16`, func(e *colly.HTMLElement) {
		nama := e.ChildText("[class=’p.product-title text-dark pt-12 non-loading’]")
		gambar := e.ChildAttr("img", "src")
		harga := e.ChildText("[class='text price text-coral text-left is-weight-bold ml-8 position-relative pb-8  non-loading'] .original-price strikethrough text-charcoal align-center med-weight non-loading")

		// Write all scraped pieces of information to output text file
		writer.Write([]string{
			nama,
			gambar,
			harga,
		})
	})

	// start scraping the page under the given URL
	c.Visit(fetchURL)
	fmt.Println("End of scraping: ", fetchURL)
}

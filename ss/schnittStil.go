// https://godoc.org/github.com/PuerkitoBio/goquery

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL = "https://www.stoffundstil.de"

func runScrape() {
	var resultCount = 1
	var page = 1
	for resultCount != 0 {
		url := baseURL + "/schnittmuster#IsPaging=false&Page=" + strconv.Itoa(page)
		fmt.Printf("Url: %v\n", url)
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		result := doc.Find("ul.product-list li.product-list-item").Each(extractPattern)
		resultCount = len(result.Nodes)
		page++
	}
}

func extractPattern(i int, s *goquery.Selection) {
	//musterTitle, _ := s.Find("h2 a").Attr("title")
	patternURL, _ := s.Find("div.productMedia-list li a").Attr("href")
	//title := s.Find("i").Text()
	//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	//fmt.Printf("Muster %d: %s\nURL: %s\n\n", i, musterTitle, patternURL)
	doc, err := goquery.NewDocument(baseURL + patternURL)
	if err != nil {
		log.Fatal(err)
	}
	patternName := doc.Find("h1[itemprop='name']").Text()
	patternPrice := doc.Find("div.product-price").Text()
	fmt.Printf("Muster %d: %s\nPreis: %s\n\n", i, patternName, patternPrice)
}

func main() {
	//exampleScrape()
	runScrape()
}

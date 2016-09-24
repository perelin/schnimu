// https://godoc.org/github.com/PuerkitoBio/goquery
/*
Crawler Config:
- baseURL
- startPath
- paginated
- patternURLs full URI or just Path
- overview list selector
- detail selectors (schema)

*/
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func exampleScrape() {
	doc, err := goquery.NewDocument("https://www.stoffe.de/schnittmuster.html?page=69")
	if err != nil {
		log.Fatal(err)
	}
	result := doc.Find("ul[data-list='Schnittmuster'] li").Each(extractPattern)
	fmt.Printf("Result: %v", len(result.Nodes))
}

func runScrape() {
	var resultCount = 1
	var page = 1
	for resultCount != 0 {
		url := "https://www.stoffe.de/schnittmuster.html?page=" + strconv.Itoa(page)
		fmt.Printf("Url: %v", url)
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		result := doc.Find("ul[data-list='Schnittmuster'] li").Each(extractPattern)
		resultCount = len(result.Nodes)
		page++
	}
}

func extractPattern(i int, s *goquery.Selection) {
	//musterTitle, _ := s.Find("h2 a").Attr("title")
	patternURL, _ := s.Find("h2 a").Attr("href")
	//title := s.Find("i").Text()
	//fmt.Printf("Review %d: %s - %s\n", i, band, title)
	//fmt.Printf("Muster %d: %s\nURL: %s\n\n", i, musterTitle, patternURL)
	doc, err := goquery.NewDocument(patternURL)
	if err != nil {
		log.Fatal(err)
	}
	patternName := doc.Find("div.article-name h1").Text()
	patternPrice := doc.Find("div.article-price div.value").Text()
	breadcrumb, _ := doc.Find("div.breadcrumb").Html()
	fmt.Printf("Muster %d: %s\nPreis: %s\n\nBreadC: %s\n\n", i, patternName, patternPrice, breadcrumb)
}

func main() {
	//exampleScrape()
	runScrape()
}

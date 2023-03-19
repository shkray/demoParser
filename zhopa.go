package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getDocument(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func parseLaptopsPrices(url string) {
	doc, err := getDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".catalog-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".catalog-item__name").Text()
		price := s.Find(".catalog-item__price-current").Text()
		fmt.Printf("Laptop %d: %s - %s\n", i+1, name, price)
	})
}

func main() {
	url := "https://www.dns-shop.ru/catalog/17a892f816404e77/noutbuki/"
	parseLaptopsPrices(url)
}

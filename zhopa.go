package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func getDocument(url string) *goquery.Document {
	http, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer http.Body.Close()
	resp, err := goquery.NewDocumentFromReader(http.Body)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func main() {
	//url := "https://www.dns-shop.ru/catalog/17a892f816404e77/noutbuki/"
	// url := "www.dns-shop.ru/ajax-state/product-buy/"

	// client := &http.Client{}

	// req, _ := http.NewRequest("POST", url, nil)

	// req.Header.Add("X-Requested-With", "XMLHttpRequest")
	// req.Header.Add("X-CSRF-Token", "702e3CCMcsYclqQunaxZEdxQY9TwXy6stjQstpoMGZvfHdqkVvVDrGrd0xrKmh5YryYRg4QuduTBf276w3pRog==")
	// req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	// req.Header.Add("Referer", "https://www.dns-shop.ru/catalog/17a892f816404e77/noutbuki/no-referrer")

	// resp, _ := client.Do(req)

	// fmt.Print(req)
	// fmt.Print(resp)

	url := "https://www.citilink.ru/catalog/noutbuki/"
	resp := getDocument(url)
	file, err := os.Create("filik.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(resp.Text()))
}

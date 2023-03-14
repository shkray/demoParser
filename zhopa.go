package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func getResponce(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //получение объекта http.Responce, обработка ошибок и закрытие чтения
	return resp
}

func main() {
	url := "https://vk.com"
	resp := getResponce(url)

	file, err := os.Create("httpTree.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //создание файла для записи и обработка ошибок

	httpResp, err := goquery.NewDocumentFromReader(resp.Body)
}

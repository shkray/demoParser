package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://vk.com"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() //получение объекта http.Responce, обработка ошибок и закрытие чтения для избежания утечек памяти

	file, err := os.Create("httpTree.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //создание файла для записи и обработка ошибок
	httpResp, err := goquery.NewDocumentFromReader(resp.Body)
}

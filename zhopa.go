package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Выполняем запрос к странице
	response, err := http.Get("https://www.example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Создаем новый объект goquery из ответа
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Извлекаем HTML-код страницы
	html, err := doc.Html()
	if err != nil {
		log.Fatal(err)
	}

	// Выводим полученный HTML-код на экран
	fmt.Println(html)
}

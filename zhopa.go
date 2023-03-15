package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://vk.com"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("httpTree.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(string(respBody))
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// func getDocument(url string) *goquery.Document {
// 	http, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer http.Body.Close()
// 	resp, err := goquery.NewDocumentFromReader(http.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return resp
// }

func clientProxy(rawProxy *string) *http.Client {
	username := (*rawProxy)[0:strings.Index(*rawProxy, ":")]
	password := (*rawProxy)[strings.Index(*rawProxy, ":")+1 : strings.Index(*rawProxy, "@")]
	proxy := "https://" + strings.Replace((*rawProxy)[strings.Index(*rawProxy, "@")+1:], "\n", "", -1)
	fmt.Print(rawProxy)
	proxyUrl, err := url.Parse(proxy)
	proxyAuth := url.UserPassword(username, password)
	if err != nil {
		log.Fatal(err)
	}
	proxyUrl.User = proxyAuth
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	return httpClient
}

func getProxyURLs() []string {
	file, err := os.Open(".proxy")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var proxyList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Print(line)
		proxyList = append(proxyList, line)
	}
	return proxyList
}

func main() {
	//url := "http://www.dns-shop.ru/catalog/17a892f816404e77/noutbuki/"
	proxyList := getProxyURLs()
	proxyNumber := 0
	//fmt.Print(proxyList[0], "\n")
	client := clientProxy(&proxyList[proxyNumber])
	fmt.Println(client)
}

//token, err := httpDoc.Find("meta[name='csrf-token']").Attr("content")   выборка csrf токена в днс

package common

import (
	"log"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Ternary(str string) (ret string) {
	if len(str) != 0 {
		ret = str
	} else {
		ret = "N/A"
	}

	return ret
}

func OnError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func GetPage(url string) (document *goquery.Document) {
	response, err := http.Get(url)
	OnError("Error getting page. ", err)

	defer response.Body.Close()

	document, err = goquery.NewDocumentFromReader(response.Body)
	OnError("Error loading body. ", err)

	return document
}

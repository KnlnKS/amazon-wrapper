package item

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func GetItemPage(item string, variant string) (document *goquery.Document) {
	response, err := http.Get("https://www.amazon.ca/" + item + "/dp/" + variant)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	return document
}

func GetItem(item string, variant string) (productTitle string) {
	document := GetItemPage(item, variant)
	productTitle = strings.TrimSpace(document.Find("#productTitle").Text())
	return productTitle
}

func Start(c *gin.Context) {
	item := c.Param("item")
	variant := c.Param("variant")

	productTitle := GetItem(item, variant)

	c.JSON(200, gin.H{
		"productTitle": productTitle,
		"item":         item,
		"variant":      variant,
	})
}

package item

import (
	"amazon-wrapper/common"

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

func GetItem(item string, variant string) (res gin.H) {
	document := GetItemPage(item, variant)

	productTitle := strings.TrimSpace(document.Find("#productTitle").Text())
	price := document.Find("#priceblock_ourprice").Text()
	rating := document.Find(".a-icon-alt").First().Text()
	numReviews := document.Find("#acrCustomerReviewText").First().Text()
	landingImage, _ := document.Find("#landingImage").Attr("src")
	hiResLandingImage, _ := document.Find("#landingImage").Attr("data-old-hires")

	res = gin.H{
		"url":               "https://www.amazon.ca/" + item + "/dp/" + variant,
		"productTitle":      common.Ternary(productTitle),
		"price":             common.Ternary(price),
		"numReviews":        common.Ternary(numReviews),
		"rating":            common.Ternary(rating),
		"landingImage":      common.Ternary(landingImage),
		"hiResLandingImage": common.Ternary(hiResLandingImage),
	}

	return res
}

func Start(c *gin.Context) {
	item := c.Param("item")
	variant := c.Param("variant")

	c.JSON(200, GetItem(item, variant))
}

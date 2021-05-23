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

func GetItem(item string, variant string) (productTitle string, price string, rating string, numReviews string, landingImage string, hiResLandingImage string) {
	document := GetItemPage(item, variant)
	var exists bool

	productTitle = strings.TrimSpace(document.Find("#productTitle").Text())
	price = document.Find("#priceblock_ourprice").Text()
	rating = document.Find(".a-icon-alt").First().Text()
	numReviews = document.Find("#acrCustomerReviewText").First().Text()
	landingImage, exists = document.Find("#landingImage").Attr("src")
	if !exists {
		log.Println("Cannot find Landing Image")
		landingImage = "N/A"
	}
	hiResLandingImage, exists = document.Find("#landingImage").Attr("data-old-hires")
	if !exists {
		log.Println("Cannot find Hi Res Landing Image")
		hiResLandingImage = "N/A"
	}

	return productTitle, price, rating, numReviews, landingImage, hiResLandingImage
}

func Start(c *gin.Context) {
	item := c.Param("item")
	variant := c.Param("variant")

	productTitle, price, rating, numReviews, landingImage, hiResLandingImage := GetItem(item, variant)

	c.JSON(200, gin.H{
		"url":               "https://www.amazon.ca/" + item + "/dp/" + variant,
		"productTitle":      productTitle,
		"price":             price,
		"numReviews":        numReviews,
		"rating":            rating,
		"landingImage":      landingImage,
		"hiResLandingImage": hiResLandingImage,
	})
}

package item

import (
	"amazon-wrapper/common"

	"strings"

	"github.com/gin-gonic/gin"
)

func getItem(item string, variant string) (res gin.H) {
	document := common.GetPage("https://www.amazon.ca/" + item + "/dp/" + variant)

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

	c.JSON(200, getItem(item, variant))
}

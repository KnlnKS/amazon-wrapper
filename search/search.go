package search

import (
	"amazon-wrapper/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func getSearchResults(searchTerm string) (res gin.H) {
	document := common.GetPage("https://www.amazon.ca/s?k=" + searchTerm)

	var parsedSearchResults []gin.H

	searchResults := document.Find(`div[data-component-type="s-search-result"]`)
	for i := range searchResults.Nodes {
		result := searchResults.Eq(i)

		sponsored := result.Find(".s-sponsored-label-text").Text() != ""
		productTitle := result.Find("h2>.a-link-normal").Text()
		rating := result.Find(".a-icon-star-small > span").First().Text()
		price := result.Find(".a-price > .a-offscreen").First().Text()
		smallImage, _ := result.Find(".s-image").Attr("src")
		imageSet, _ := result.Find(".s-image").Attr("srcset")
		url, _ := result.Find(".s-line-clamp-4>.a-link-normal").Attr("href")

		endpoint := ""

		if !sponsored {
			endpoint = "item/" + strings.Split(url, "/")[1] + "/" + strings.Split(url, "/")[3]
		}

		parsedSearchResults = append(parsedSearchResults, gin.H{
			"productTitle": productTitle,
			"rating":       rating,
			"price":        price,
			"smallImage":   smallImage,
			"imageSet":     strings.Split(imageSet, ", "),
			"sponsored":    sponsored,
			"url":          url,
			"endpoint":     endpoint,
		})

	}

	res = gin.H{
		"searchTerm":    searchTerm,
		"searchResults": parsedSearchResults,
	}

	return res
}

func Start(c *gin.Context) {
	searchTerm := c.Query("s")

	c.JSON(200, getSearchResults(searchTerm))
}

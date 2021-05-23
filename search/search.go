package search

import (
	"amazon-wrapper/common"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func extractSearchResult(i int, result *goquery.Selection) {
	fmt.Println(result.Text())
}

func getSearchResults(searchTerm string) (res gin.H) {
	document := common.GetPage("https://www.amazon.ca/s?k=" + searchTerm)

	document.Find(`div[data-component-type="s-search-result"]`).Each(extractSearchResult)

	res = gin.H{
		"searchTerm": searchTerm,
	}

	return res
}

func Start(c *gin.Context) {
	searchTerm := c.Query("s")
	getSearchResults(searchTerm)
	c.JSON(200, gin.H{
		"searchTerm": searchTerm,
	})
}

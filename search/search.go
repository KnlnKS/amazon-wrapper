package search

import (
	"amazon-wrapper/common"
	"fmt"

	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func getSearchResultsPage(searchTerm string) (document *goquery.Document) {
	response, err := http.Get("https://www.amazon.ca/s?k=" + searchTerm)
	common.OnError("Error getting page. ", err)

	defer response.Body.Close()

	document, err = goquery.NewDocumentFromReader(response.Body)
	common.OnError("Error loading body. ", err)

	return document
}

func extractSearchResult(i int, result *goquery.Selection) {
	fmt.Println(result.Text())
}

func getSearchResults(searchTerm string) (res gin.H) {
	document := getSearchResultsPage(searchTerm)

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

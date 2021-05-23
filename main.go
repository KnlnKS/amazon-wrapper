package main

import (
	"amazon-wrapper/item"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/item/:item/:variant", item.Start)

	r.GET("/search", func(c *gin.Context) {
		searchTerm := c.Query("s")
		c.JSON(200, gin.H{
			"searchTerm": searchTerm,
		})
	})

	r.Run()
}

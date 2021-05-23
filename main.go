package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/item/:item", func(c *gin.Context) {
		item := c.Param("item")
		c.JSON(200, gin.H{
			"item": item,
		})
	})

	r.GET("/search", func(c *gin.Context) {
		searchTerm := c.Query("s")
		c.JSON(200, gin.H{
			"searchTerm": searchTerm,
		})
	})

	r.Run()
}

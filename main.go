package main

import (
	"amazon-wrapper/item"
	"amazon-wrapper/search"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/item/:item/:variant", item.Start)
	r.GET("/search", search.Start)
	r.Run()
}

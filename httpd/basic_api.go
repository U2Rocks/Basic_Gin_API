package main

import (
	"FIRST_GIN_GONIC_API/httpd/handler"
	"FIRST_GIN_GONIC_API/platform/itemfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	// instantiate default object
	r := gin.Default()

	// create new Cvsitem list
	feed := itemfeed.New()

	// routes
	r.GET("/cvslist", handler.ItemsGet(feed))
	r.POST("/cvslist", handler.AddItem(feed))

	r.Run(":9090") // api now runs on port 9090
}

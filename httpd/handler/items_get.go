package handler

import (
	"FIRST_GIN_GONIC_API/platform/itemfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ItemsGet(feed *itemfeed.Itemlist) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetItemsBought()
		c.JSON(http.StatusOK, results)
	}
}

package handler

import (
	"FIRST_GIN_GONIC_API/platform/itemfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addItemRequest struct {
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
	Date string  `json:"date"`
}

func AddItem(feed *itemfeed.Itemlist) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := addItemRequest{}
		c.Bind(&requestBody)

		item := itemfeed.Cvsitem{
			Name: requestBody.Name,
			Cost: requestBody.Cost,
			Date: requestBody.Date,
		}

		feed.AddItemtoList(item)

		c.Status(http.StatusNoContent)
	}
}

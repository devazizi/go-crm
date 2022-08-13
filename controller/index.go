package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeAPI() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"url":     c.FullPath(),
			"content": "not found",
		})
	}
}

package controller

import (
	"github.com/devazizi/go-crm/infrastructure"
	"github.com/gin-gonic/gin"
)

func IndexTasks(DB infrastructure.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": true})
	}
}

package controller

import (
	"github.com/devazizi/go-crm/contract/response"
	"github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexTaskCategory(DB infrastructure.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskCategories := repository.New(DB).IndexTaskCategory()

		c.JSON(http.StatusOK, response.Response{
			Status: true,
			Data:   taskCategories,
		})
	}
}

func GetTaskCategory(DB infrastructure.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskCategoryId, _ := strconv.Atoi(c.Param("taskCategoryId"))

		taskCategory, err := repository.New(DB).GetTaskCategory(taskCategoryId)

		if err != nil {
			c.JSON(http.StatusNotFound, response.Response{Message: "not found"})
			return
		}

		c.JSON(http.StatusOK, response.Response{Status: true, Data: taskCategory})
	}
}

func DeleteTaskCategory(DB infrastructure.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskCategoryId, _ := strconv.Atoi(c.Param("taskCategoryId"))

		taskCategory, err := repository.New(DB).GetTaskCategory(taskCategoryId)

		if err != nil {
			c.JSON(http.StatusNotFound, response.Response{Message: "not found"})
			return
		}

		repository.New(DB).DeleteTaskCategory(taskCategory)

		c.JSON(http.StatusOK, response.Response{Status: true})
	}
}

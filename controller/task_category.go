package controller

import (
	"github.com/devazizi/go-crm/contract"
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/contract/response"
	"github.com/devazizi/go-crm/entity"
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

func CreateTaskCategory(DB infrastructure.DB, validator contract.ValidateCreateTaskCategoryRequest) gin.HandlerFunc {

	return func(c *gin.Context) {
		createTaskCategoryRequest := request.CreateTaskCategoryRequest{}

		if err := c.BindJSON(&createTaskCategoryRequest); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: "Bad Request"})
			return
		}

		if err := validator(createTaskCategoryRequest); err != nil {
			c.JSON(http.StatusUnprocessableEntity, response.Response{Message: err.Error()})
			return
		}

		taskCategory := entity.TaskCategory{
			Name: createTaskCategoryRequest.Name,
		}

		data := repository.New(DB).CreateTaskCategory(taskCategory)

		c.JSON(http.StatusOK, response.Response{Status: true, Data: data})
	}

}

func UpdateTaskCategory(DB infrastructure.DB, validator contract.ValidateUpdateTaskCategoryRequest) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskCategoryId, _ := strconv.Atoi(c.Param("taskCategoryId"))
		updateTaskCategoryRequest := request.UpdateTaskCategoryRequest{}

		if err := c.BindJSON(&updateTaskCategoryRequest); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: "Bad Request"})
			return
		}

		if err := validator(updateTaskCategoryRequest); err != nil {
			c.JSON(http.StatusUnprocessableEntity, response.Response{Message: err.Error()})
			return
		}

		taskCategory, notfound := repository.New(DB).GetTaskCategory(taskCategoryId)

		if notfound != nil {
			if err := c.BindJSON(&updateTaskCategoryRequest); err != nil {
				c.JSON(http.StatusNotFound, response.Response{Message: "Not found"})
				return
			}
		}

		taskCategory.Name = updateTaskCategoryRequest.Name

		repository.New(DB).UpdateTaskCategory(taskCategory)

		c.JSON(http.StatusOK, response.Response{Status: true, Data: taskCategory})
	}

}

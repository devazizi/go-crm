package controller

import (
	"github.com/devazizi/go-crm/contract"
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/contract/response"
	"github.com/devazizi/go-crm/entity"
	"github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/repository"
	"github.com/devazizi/go-crm/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexTasks(DB infrastructure.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks := repository.New(DB).IndexTask()

		c.JSON(http.StatusOK, response.Response{
			Status: true,
			Data:   tasks,
		})
	}
}

func GetTask(DB infrastructure.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskId, _ := strconv.Atoi(c.Param("taskId"))

		task, err := repository.New(DB).GetTask(taskId)

		if err != nil {
			c.JSON(http.StatusNotFound, response.Response{Message: "not found"})
			return
		}

		c.JSON(http.StatusOK, response.Response{Status: true, Data: task})
	}
}

func CreateTask(DB infrastructure.DB, validator contract.ValidateCreateTaskRequest) gin.HandlerFunc {

	return func(c *gin.Context) {
		createTaskRequest := request.CreateTaskRequest{}
		if err := c.BindJSON(&createTaskRequest); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: "Bad Request"})
			return
		}

		if err := validator(createTaskRequest); err != nil {
			c.JSON(http.StatusUnprocessableEntity, response.Response{Message: err.Error()})
			return
		}

		auth := auth.GetAuthentication()

		var assignTo []entity.User

		for _, userId := range createTaskRequest.AssignTo {
			assignTo = append(assignTo, entity.User{ID: uint(userId)})
		}

		task := entity.Task{
			User:        entity.User{ID: uint(auth.UserId)},
			Name:        createTaskRequest.Name,
			CategoryId:  createTaskRequest.CategoryID,
			AssignTo:    assignTo,
			Description: createTaskRequest.Description,
		}

		repository.New(DB).CreateTask(task)

		c.JSON(http.StatusCreated, createTaskRequest)
	}
}

func DeleteTask(DB infrastructure.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		taskId, _ := strconv.Atoi(c.Param("taskId"))

		task, err := repository.New(DB).GetTask(taskId)

		if err != nil {
			c.JSON(http.StatusNotFound, response.Response{Message: "not found"})
			return
		}

		repository.New(DB).DeleteTask(task)

		c.JSON(http.StatusOK, response.Response{Status: true, Data: task})
	}
}

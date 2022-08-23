package router

import (
	"github.com/devazizi/go-crm/contract/validation"
	"github.com/devazizi/go-crm/controller"
	infra "github.com/devazizi/go-crm/infrastructure"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, database infra.DB, redis infra.RedisConnection) {

	router.Use(middlewareCORS())
	router.GET("/", controller.HomeAPI())

	apiV1 := router.Group("/api/v1")
	{
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/login", controller.LoginAPI(database, validation.ValidateLoginRequestFields))
			authRoutes.POST("/register", controller.RegisterAPI(database, validation.ValidateRegisterRequestFields))
			authRoutes.POST("/forget-password", controller.ForgetPassword(database, validation.ValidateForgetPasswordRequestFields))
			authRoutes.Use(middlewareCheckAuthenticated(database)).POST(
				"/change-password",
				controller.ChangePassword(database, validation.ValidateChangePasswordRequestFields),
			)
		}

		taskRoutes := apiV1.Group("/tasks").Use(middlewareCheckAuthenticated(database))
		{
			taskRoutes.GET("/", controller.IndexTasks(database))
			taskRoutes.POST("/", controller.CreateTask(database, validation.ValidateCreateTaskRequest))
			taskRoutes.GET("/:taskId", controller.GetTask(database))
			taskRoutes.DELETE("/:taskId", controller.DeleteTask(database))
		}

		taskCategoryRoutes := apiV1.Group("/task-categories").Use(middlewareCheckAuthenticated(database))
		{
			taskCategoryRoutes.GET("/", controller.IndexTaskCategory(database))
			taskCategoryRoutes.POST("/", controller.CreateTaskCategory(database, validation.ValidateCreateTaskCategoryRequest))
			taskCategoryRoutes.GET("/:taskCategoryId", controller.GetTaskCategory(database))
			taskCategoryRoutes.DELETE("/:taskCategoryId", controller.DeleteTaskCategory(database))
			taskCategoryRoutes.PUT("/:taskCategoryId", controller.UpdateTaskCategory(database, validation.ValidateUpdateTaskCategoryRequest))
		}
	}
}

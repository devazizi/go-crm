package main

import (
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/controller"
	infra "github.com/devazizi/go-crm/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {

	dsn := "host=localhost user=postgres password=12345678 dbname=go_01 port=5432 sslmode=disable TimeZone=Asia/Tehran"
	DbConnection := infra.NewDB(dsn)

	routerEngine := gin.Default()
	router(routerEngine, DbConnection)

	routerEngine.Run(":9000")
}

func router(router *gin.Engine, database infra.DB) {
	router.GET("/", controller.HomeAPI())

	apiV1 := router.Group("/api/v1")
	{
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/login", controller.LoginAPI(database, request.LoginRequest{}))
			authRoutes.POST("/register", controller.RegisterAPI(database))
			//authRoutes.POST("/forget-password", controller.ForgetPassword(nil))
		}
	}

	//usersRoutes := router.Group("/users")
	//{
	//	//usersRoutes.POST("/", controller.)
	//	//usersRoutes.GET("/", controller)
	//	//usersRoutes.GET("/:userId")
	//	//usersRoutes.PUT("/:userId")
	//	//usersRoutes.DELETE("/:userId")
	//	//usersRoutes.GET("/user")
	//}
	//
	//taskRoutes := router.Group("/tasks")
	//{
	//	//taskRoutes.GET("/")
	//	//taskRoutes.GET("/:taskId")
	//	//taskRoutes.DELETE("/:taskId")
	//}
}

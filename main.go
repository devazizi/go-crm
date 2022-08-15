package main

import (
	"fmt"
	"github.com/devazizi/go-crm/contract/response"
	"github.com/devazizi/go-crm/contract/validation"
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
			authRoutes.POST("/login", controller.LoginAPI(database, validation.ValidateLoginRequestFields))
			authRoutes.POST("/register", controller.RegisterAPI(database, validation.ValidateRegisterRequestFields))
			//authRoutes.POST("/forget-password", controller.ForgetPassword(nil))
		}
		taskRoutes := apiV1.Group("/tasks").Use(middlewareCheckAuthenticated())
		{
			taskRoutes.GET("/", controller.IndexTasks(database))
			//	//taskRoutes.GET("/:taskId")
			//	//taskRoutes.DELETE("/:taskId")
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

}

func middlewareCheckAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationKey := c.GetHeader("Authorization")

		if authorizationKey == "" {
			c.JSON(401, response.Response{Status: false, Message: "authorization fail"})
			return
		}

		fmt.Println(authorizationKey)

		c.Next()
	}
}

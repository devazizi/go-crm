package main

import (
	"fmt"
	"github.com/devazizi/go-crm/controller"
	"github.com/devazizi/go-crm/entity"
	"github.com/devazizi/go-crm/service/jwt"
	"github.com/gin-gonic/gin"
)

func main() {

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//DbConnection := infra.NewDB(dsn)

	user := entity.User{Name: "alireza aizizi", ID: 433434, Email: "dev.azizimail@gmail.com"}
	token, _ := jwt.CreateToken("2332-abcd-acdf-ccd2", user)
	fmt.Println(token)
	routerEngine := gin.Default()
	//router(routerEngine, DbConnection)
	router(routerEngine)
	routerEngine.Run(":9000")
}

//func router(router *gin.Engine, Database infra.DB) {
//	router.GET("/", controller.HomeAPI())
//}

func router(router *gin.Engine) {
	router.GET("/", controller.HomeAPI())

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", controller.LoginAPI(nil))
		authRoutes.POST("/register", controller.RegisterAPI(nil))
		authRoutes.POST("/forget-password", controller.ForgetPassword(nil))
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

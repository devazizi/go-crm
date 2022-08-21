package main

import (
	"net/http"
	"strings"

	"github.com/devazizi/go-crm/contract/response"
	"github.com/devazizi/go-crm/contract/validation"
	"github.com/devazizi/go-crm/controller"
	infra "github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/repository"
	"github.com/devazizi/go-crm/service/auth"
	"github.com/devazizi/go-crm/service/jwt"
	"github.com/gin-gonic/gin"
)

func main() {

	databaseDsn := "host=localhost user=postgres password=12345678 dbname=go_01 port=5432 sslmode=disable TimeZone=Asia/Tehran"
	redisDsn := map[string]any{
		"addr":     "localhost:6379",
		"password": "",
		"db":       0,
	}
	DbConnection := infra.NewDB(databaseDsn)
	RedisConnection := infra.NewRedis(redisDsn)
	routerEngine := gin.Default()
	router(routerEngine, DbConnection, RedisConnection)

	routerEngine.Run("localhost:9000")
}

func router(router *gin.Engine, database infra.DB, redis infra.RedisConnection) {

	router.Use(middlewareCros())
	router.GET("/", controller.HomeAPI())

	apiV1 := router.Group("/api/v1")
	{
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/login", controller.LoginAPI(database, validation.ValidateLoginRequestFields))
			authRoutes.POST("/register", controller.RegisterAPI(database, validation.ValidateRegisterRequestFields))
			//authRoutes.POST("/forget-password", controller.ForgetPassword(nil))
		}
		taskRoutes := apiV1.Group("/tasks").Use(middlewareCheckAuthenticated(database))

		{
			taskRoutes.GET("/", controller.IndexTasks(database))
			taskRoutes.POST("/", controller.CreateTask(database, validation.ValidateCreateTaskRequest))
			taskRoutes.GET("/:taskId", controller.GetTask(database))
			taskRoutes.DELETE("/:taskId", controller.DeleteTask(database))
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
func middlewareCros() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}
}

func middlewareCheckAuthenticated(DB infra.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationKey := c.GetHeader("Authorization")

		if authorizationKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{Status: false, Message: "authorization fail"})
			return
		}

		tokenData := strings.Split(authorizationKey, " ")

		jwtCredentials, _ := jwt.GetClaimsFromToken(tokenData[1])
		userInfo := jwtCredentials["UserInfo"]
		user := userInfo.(map[string]interface{})

		userId := user["id"].(float64)
		tokenHash, isCorrect := user["token"].(string)

		if !isCorrect {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{Status: false, Message: "authentication fail"})
			return
		}

		isValid := repository.ValidateTokenExistInStorage(DB, tokenHash, int(userId))

		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{Status: false, Message: "authentication fail"})
			return
		}

		authentication := auth.Authentication{UserId: int(userId), Token: tokenData[1]}
		authentication.SetAuthentication()

		c.Next()
	}
}

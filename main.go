package main

import (
	"github.com/devazizi/go-crm/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//DbConnection := infra.NewDB(dsn)

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
}

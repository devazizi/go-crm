package main

import (
	infra "github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func environmentVariables() {
	godotenv.Load(".env")
}

func prepareRedis() map[string]any {

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return map[string]any{
		"addr":     os.Getenv("REDIS_ADDRESS"),
		"password": os.Getenv("REDIS_PASSWORD"),
		"db":       redisDB,
	}
}

func main() {
	environmentVariables()

	databaseDsn := os.Getenv("DATABASE_DSN")

	DbConnection := infra.NewDB(databaseDsn)
	RedisConnection := infra.NewRedis(prepareRedis())
	routerEngine := gin.Default()
	router.Router(routerEngine, DbConnection, RedisConnection)

	routerEngine.Run("localhost:9000")
}

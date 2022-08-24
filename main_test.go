package main

import (
	infra "github.com/devazizi/go-crm/infrastructure"
	"os"
	"reflect"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	environmentVariables()

	databaseDsn := os.Getenv("DATABASE_DSN")
	dbConnection := infra.NewDB(databaseDsn)

	Connected := reflect.TypeOf(dbConnection) == reflect.TypeOf(infra.DB{})
	if !Connected {
		t.Error("can not connect to db")
	}
}

func TestRedisConnection(t *testing.T) {
	environmentVariables()
	redisConnectionInfo := prepareRedis()
	redisConnection := infra.NewRedis(redisConnectionInfo)

	connected := reflect.TypeOf(redisConnection) == reflect.TypeOf(infra.RedisConnection{})

	if !connected {
		t.Error(connected)
	}
}

package infrastructure

import (
	"github.com/devazizi/go-crm/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB(dsn string) DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can not connect to database")
	}

	var entities = []any{
		&entity.User{},
		&entity.Task{},
		&entity.TaskCategory{},
		&entity.Token{},
	}

	if autoMigrate := db.AutoMigrate(entities...); autoMigrate != nil {
		panic("can not auto migrate")
	}

	return DB{Connection: db}
}

package repository

import (
	"github.com/devazizi/go-crm/entity"
	infra "github.com/devazizi/go-crm/infrastructure"
)

type UserRepository interface {
	CreateUser(DB infra.DB, user entity.User)
}

func CreateUser(DB infra.DB, user entity.User) {
	DB.Connection.Create(&user)
}

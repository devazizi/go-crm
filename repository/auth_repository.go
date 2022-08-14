package repository

import (
	"github.com/devazizi/go-crm/entity"
	infra "github.com/devazizi/go-crm/infrastructure"
)

type AuthRepository interface {
	CheckUserExistsByEmail(DB infra.DB, email string) (entity.User, error)
}

func CheckUserExistsByEmail(DB infra.DB, email string) (entity.User, error) {
	var user entity.User

	result := DB.Connection.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

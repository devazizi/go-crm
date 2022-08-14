package repository

import (
	"github.com/devazizi/go-crm/entity"
	infra "github.com/devazizi/go-crm/infrastructure"
)

type AuthRepository interface {
	CheckUserExistsByEmail(DB infra.DB, email string) (entity.User, error)
	CheckEmailIsUnique(DB infra.DB, email string) bool
}

func CheckUserExistsByEmail(DB infra.DB, email string) (entity.User, error) {
	var user entity.User

	result := DB.Connection.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func CheckEmailIsUnique(DB infra.DB, email string) bool {
	var user entity.User
	DB.Connection.Model(user).Where("email = ?", email).Find(&user)

	if user.ID == 0 {
		return true
	}

	return false
}

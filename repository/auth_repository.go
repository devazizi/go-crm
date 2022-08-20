package repository

import (
	"github.com/devazizi/go-crm/entity"
	infra "github.com/devazizi/go-crm/infrastructure"
	"time"
)

type AuthRepository interface {
	CheckUserExistsByEmail(DB infra.DB, email string) (entity.User, error)
	CheckEmailIsUnique(DB infra.DB, email string) bool
	CreateClientToken(DB infra.DB, user entity.User, tokenHash string) (entity.Token, error)
	ValidateTokenExistInStorage(DB infra.DB, token string, userId int) bool
}

func CreateClientToken(DB infra.DB, user entity.User, tokenHash string) (entity.Token, error) {

	token := entity.Token{UserID: int(user.ID), Hash: tokenHash, ExpiryDate: time.Now().AddDate(0, 0, 30)}
	result := DB.Connection.Create(&token)

	if result.Error != nil {
		return token, result.Error
	}

	return token, nil
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

func ValidateTokenExistInStorage(DB infra.DB, tokenHash string, userId int) bool {
	var token entity.Token
	err := DB.Connection.Where("hash = ?", tokenHash).Where("user_id = ?", userId).First(&token)

	if err.Error != nil {
		return false
	}

	return true
}

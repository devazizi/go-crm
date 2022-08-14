package entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name", gorm:"size:256"`
	Email     string    `json:"email" gorm:"size:256"`
	password  string    `gorm:"size:512"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) SetPassword(newPassword string) bool {
	user.password = newPassword

	return true
}

func (user *User) GetPassword() string {
	return user.password
}

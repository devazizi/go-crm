package entity

import "time"

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	password  string
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

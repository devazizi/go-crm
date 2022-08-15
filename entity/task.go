package entity

import "time"

type Task struct {
	ID          uint      `json:"id"`
	CategoryId  int       `json:"category_id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	User        User      `json:"user"`
	AssignTo    []User    `gorm:"many2many:user_task;" gorm:"references:UserID" json:"assign_to"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

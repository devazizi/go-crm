package repository

import (
	"github.com/devazizi/go-crm/entity"
	"github.com/devazizi/go-crm/infrastructure"
)

type TaskRepository interface {
	CreateTask(DB infrastructure.DB, task entity.Task)
}

func CreateTask(DB infrastructure.DB, task entity.Task) {
	DB.Connection.Create(&task)
	DB.Connection.Save(&task)
}

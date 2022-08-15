package repository

import (
	"github.com/devazizi/go-crm/entity"
	"github.com/devazizi/go-crm/infrastructure"
)

type TaskRepository interface {
	CreateTask(DB infrastructure.DB, task entity.Task)
	IndexTask(DB infrastructure.DB) []TaskRepository
}

func CreateTask(DB infrastructure.DB, task entity.Task) {
	DB.Connection.Create(&task)
	DB.Connection.Save(&task)
}

func IndexTask(DB infrastructure.DB) []entity.Task {
	var tasks []entity.Task

	DB.Connection.Preload("AssignTo").Find(&tasks)

	return tasks
}

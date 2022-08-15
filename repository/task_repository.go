package repository

import (
	"errors"
	"github.com/devazizi/go-crm/entity"
	"github.com/devazizi/go-crm/infrastructure"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(DB infrastructure.DB, task entity.Task)
	IndexTask(DB infrastructure.DB) []TaskRepository
	GetTask(DB infrastructure.DB, taskId int) (entity.Task, error)
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

func GetTask(DB infrastructure.DB, taskId int) (entity.Task, error) {
	var task entity.Task

	result := DB.Connection.Preload("AssignTo").Preload("User").First(&task, taskId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return task, result.Error
	}

	return task, nil
}

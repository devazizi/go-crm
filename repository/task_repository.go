package repository

import (
	"errors"
	"github.com/devazizi/go-crm/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task entity.Task)
	IndexTask() []TaskRepository
	GetTask(taskId int) (entity.Task, error)
	DeleteTask(task entity.Task) error
}

func (i Interactor) CreateTask(task entity.Task) {
	i.store.Create(&task)
}

func (i Interactor) DeleteTask(task entity.Task) error {
	err := i.store.Model(&task).Association("AssignTo").Clear()
	if err != nil {
		return err
	}
	i.store.Delete(&task)

	return nil
}

func (i Interactor) IndexTask() []entity.Task {
	var tasks []entity.Task

	i.store.Preload("AssignTo").Find(&tasks)

	return tasks
}

func (i Interactor) GetTask(taskId int) (entity.Task, error) {
	var task entity.Task

	result := i.store.Preload("AssignTo").Preload("User").First(&task, taskId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return task, result.Error
	}

	return task, nil
}

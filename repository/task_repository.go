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

func (intractor Interactor) CreateTask(task entity.Task) {
	intractor.store.Create(&task)
	intractor.store.Save(&task)
}

func (intractor Interactor) DeleteTask(task entity.Task) error {
	err := intractor.store.Model(&task).Association("AssignTo").Clear()
	if err != nil {
		return err
	}
	intractor.store.Delete(&task)

	return nil
}

func (intractor Interactor) IndexTask() []entity.Task {
	var tasks []entity.Task

	intractor.store.Preload("AssignTo").Find(&tasks)

	return tasks
}

func (intractor Interactor) GetTask(taskId int) (entity.Task, error) {
	var task entity.Task

	result := intractor.store.Preload("AssignTo").Preload("User").First(&task, taskId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return task, result.Error
	}

	return task, nil
}

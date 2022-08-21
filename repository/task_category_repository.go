package repository

import (
	"errors"
	"github.com/devazizi/go-crm/entity"
	"gorm.io/gorm"
)

type TaskCategoryRepository interface {
	IndexTaskCategory() []entity.TaskCategory
	GetTaskCategory(taskCategoryId int) (entity.TaskCategory, error)
}

func (i Interactor) IndexTaskCategory() []entity.TaskCategory {
	var taskCategories []entity.TaskCategory
	i.store.Find(&taskCategories)

	return taskCategories
}

func (i Interactor) GetTaskCategory(taskCategoryId int) (entity.TaskCategory, error) {
	var taskCategory entity.TaskCategory
	result := i.store.First(&taskCategory, taskCategoryId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return taskCategory, result.Error
	}

	return taskCategory, nil
}

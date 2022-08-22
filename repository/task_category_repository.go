package repository

import (
	"errors"
	"github.com/devazizi/go-crm/entity"
	"gorm.io/gorm"
)

type TaskCategoryRepository interface {
	IndexTaskCategory() []entity.TaskCategory
	GetTaskCategory(taskCategoryId int) (entity.TaskCategory, error)
	DeleteTaskCategory(taskCategory entity.TaskCategory)
	CreateTaskCategory(taskCategory entity.TaskCategory) entity.TaskCategory
	UpdateTaskCategory(taskCategory entity.TaskCategory)
}

func (i Interactor) IndexTaskCategory() []entity.TaskCategory {
	var taskCategories []entity.TaskCategory
	i.store.Order("id desc").Find(&taskCategories)

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

func (i Interactor) DeleteTaskCategory(taskCategory entity.TaskCategory) {
	i.store.Delete(&taskCategory)
}

func (i Interactor) CreateTaskCategory(taskCategory entity.TaskCategory) entity.TaskCategory {
	i.store.Create(&taskCategory)

	return taskCategory
}

func (i Interactor) UpdateTaskCategory(taskCategory entity.TaskCategory) {
	i.store.Save(&taskCategory)
}

package repository

import (
	"github.com/devazizi/go-crm/infrastructure"
	"gorm.io/gorm"
)

type Interactor struct {
	store *gorm.DB
}

func New(db infrastructure.DB) Interactor {
	return Interactor{store: db.Connection}
}

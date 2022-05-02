package repository

import (
	"gorm.io/gorm"
)

type RepositoryBase struct {
	DB *gorm.DB
}

func NewRepositoryBase(db gorm.DB) *RepositoryBase {
	return &RepositoryBase{DB: &db}
}

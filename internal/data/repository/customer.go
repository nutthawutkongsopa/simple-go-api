package repository

import (
	"test-api/internal/data/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	*RepositoryBase
}

func NewCustomerRepository(db gorm.DB) *CustomerRepository {
	result := CustomerRepository{RepositoryBase: &RepositoryBase{}}
	result.DB = &db
	return &result
}

func (r *CustomerRepository) Get(id uuid.UUID) *entity.Customer {
	var result entity.Customer
	r.DB.Table("customer").Find(&result, id)
	return &result
}

func (r *CustomerRepository) GetAll() *gorm.DB {
	return r.DB.Table("customer")
}

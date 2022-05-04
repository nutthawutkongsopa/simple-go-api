package repository

import (
	"test-api/internal/data/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	RepositoryBase
}

func NewCustomerRepository(db gorm.DB) *CustomerRepository {
	result := new(CustomerRepository)
	result.DB = &db
	return result
}

func (r *CustomerRepository) Get(id uuid.UUID) *entity.Customer {
	var result entity.Customer
	r.DB.Table("customer").Find(&result, id)
	return &result
}

func (r *CustomerRepository) GetAll() *gorm.DB {
	return r.DB.Table("customer")
}

func (r *CustomerRepository) Save(e entity.Customer) {
	r.DB.Table("customer").Save(e)
}

func (r *CustomerRepository) Remove(e entity.Customer) {
	r.DB.Table("customer").Delete(&e)
}

package repository

import (
	"test-api/internal/core"
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

func (r *CustomerRepository) Get(id uuid.UUID) (*entity.Customer, error) {
	var result entity.Customer
	_, err := core.HandleDBError(r.DB.Table("customer").Find(&result, id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *CustomerRepository) GetAll() *gorm.DB {
	return r.DB.Table("customer")
}

func (r *CustomerRepository) Save(e entity.Customer) error {
	_, err := core.HandleDBError(r.DB.Table("customer").Save(e))
	return err
}

func (r *CustomerRepository) Remove(e entity.Customer) error {
	_, err := core.HandleDBError(r.DB.Table("customer").Delete(&e))
	return err
}

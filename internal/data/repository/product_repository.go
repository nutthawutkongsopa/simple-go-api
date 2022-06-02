package repository

import (
	"test-api/internal/core"
	"test-api/internal/data/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	RepositoryBase
}

func NewProductRepository(db gorm.DB) *ProductRepository {
	result := new(ProductRepository)
	result.DB = &db
	return result
}

func (r *ProductRepository) Get(id uuid.UUID) (*entity.Product, error) {
	var result entity.Product
	_, err := core.HandleDBError(r.DB.Table("product").Find(&result, id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ProductRepository) GetAll() *gorm.DB {
	return r.DB.Table("product")
}

func (r *ProductRepository) Save(e entity.Product) error {
	_, err := core.HandleDBError(r.DB.Table("product").Save(e))
	return err
}

func (r *ProductRepository) Remove(e entity.Product) error {
	_, err := core.HandleDBError(r.DB.Table("product").Delete(&e))
	return err
}

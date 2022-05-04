package repository

import (
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

func (r *ProductRepository) Get(id uuid.UUID) *entity.Product {
	var result entity.Product
	r.DB.Table("product").Find(&result, id)
	return &result
}

func (r *ProductRepository) GetAll() *gorm.DB {
	return r.DB.Table("product")
}

func (r *ProductRepository) Save(e entity.Product) {
	r.DB.Table("product").Save(e)
}

func (r *ProductRepository) Remove(e entity.Product) {
	r.DB.Table("product").Delete(&e)
}

package service

import (
	"test-api/internal/core"
	"test-api/internal/data/entity"
	"test-api/internal/data/repository"
	"test-api/internal/model"
	"time"

	"github.com/golobby/container/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct {
	Container container.Container
}

func NewProductService(container container.Container) *ProductService {
	sv := new(ProductService)
	sv.Container = container
	return sv
}

func (sv ProductService) Search(criteria model.ProductSearchRequest) *model.SearchResult[model.ProductDataReponse] {
	repo := core.Resolve[repository.ProductRepository](sv.Container)

	result := model.SearchResult[model.ProductDataReponse]{}
	query := sv.FilterProduct(*repo.GetAll(), criteria)
	var itemCount int64
	query.Count(&itemCount)
	result.ItemCount = int(itemCount)

	query = core.Sort(query, criteria.SortBy, criteria.SortDirection)
	resultData := core.MapQuery(query, func(x entity.Product) model.ProductDataReponse {
		return model.ProductDataReponse{
			ID:     x.ID,
			Name:   x.Name,
			Remark: x.Remark,
		}
	})
	result.ResultData = resultData

	return &result
}

func (sv ProductService) AddProduct(data model.ProductUpdateRequest) uuid.UUID {
	repo := core.Resolve[repository.ProductRepository](sv.Container)
	product := new(entity.Product)
	product.ID = uuid.New()
	product.Name = data.Name
	product.Remark = data.Remark
	product.CreatedBy = core.Ptr("Testt")
	product.CreatedDate = core.Ptr(time.Now())
	repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Save(*product)
		return nil
	})
	return product.ID
}

func (sv ProductService) UpdateProduct(id uuid.UUID, data model.ProductUpdateRequest) {
	repo := core.Resolve[repository.ProductRepository](sv.Container)
	product := repo.Get(id)
	product.Name = data.Name
	product.Remark = data.Remark
	product.UpdatedBy = core.Ptr("Testt")
	product.UpdatedDate = core.Ptr(time.Now())
	repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Save(*product)
		return nil
	})
}

func (sv ProductService) RemoveProduct(id uuid.UUID) {
	repo := core.Resolve[repository.ProductRepository](sv.Container)
	product := repo.Get(id)
	repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Remove(*product)
		return nil
	})
}

func (sv ProductService) FilterProduct(db gorm.DB, criteria model.ProductSearchRequest) gorm.DB {
	if criteria.Name != nil && *criteria.Name != "" {
		db = *db.Where("first_name like ? or last_name like ?", *criteria.Name, *criteria.Name)
	}
	return db
}

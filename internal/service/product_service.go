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

func (sv ProductService) Search(criteria model.ProductSearchRequest) (*model.SearchResult[model.ProductDataReponse], error) {
	repo, err := core.Resolve[repository.ProductRepository](sv.Container)
	if err != nil {
		return nil, err
	}

	result := model.SearchResult[model.ProductDataReponse]{}
	query := sv.FilterProduct(*repo.GetAll(), criteria)
	var itemCount int64
	core.HandleDBError(query.Count(&itemCount))
	result.ItemCount = int(itemCount)

	query = core.Sort(query, criteria.SortBy, criteria.SortDirection)
	query = core.Page(query, criteria.PageLength, criteria.Page)
	resultData, err := core.MapQuery(query, func(x entity.Product) model.ProductDataReponse {
		return model.ProductDataReponse{
			ID:     x.ID,
			Name:   x.Name,
			Remark: x.Remark,
		}
	})
	if err != nil {
		return nil, err
	}
	result.ResultData = resultData

	return &result, nil
}

func (sv ProductService) AddProduct(data model.ProductUpdateRequest) (*uuid.UUID, error) {
	repo, err := core.Resolve[repository.ProductRepository](sv.Container)
	if err != nil {
		return nil, err
	}
	product := new(entity.Product)
	product.ID = uuid.New()
	product.Name = data.Name
	product.Remark = data.Remark
	product.CreatedBy = core.Ptr("Testt")
	product.CreatedDate = core.Ptr(time.Now())
	err = repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Save(*product)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &product.ID, nil
}

func (sv ProductService) UpdateProduct(id uuid.UUID, data model.ProductUpdateRequest) error {
	repo, err := core.Resolve[repository.ProductRepository](sv.Container)
	if err != nil {
		return err
	}
	product, err := repo.Get(id)
	if err != nil {
		return err
	}
	product.Name = data.Name
	product.Remark = data.Remark
	product.UpdatedBy = core.Ptr("Testt")
	product.UpdatedDate = core.Ptr(time.Now())
	err = repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Save(*product)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (sv ProductService) RemoveProduct(id uuid.UUID) error {
	repo, err := core.Resolve[repository.ProductRepository](sv.Container)
	if err != nil {
		return err
	}
	product, err := repo.Get(id)
	if err != nil {
		return err
	}
	err = repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Remove(*product)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (sv ProductService) FilterProduct(db gorm.DB, criteria model.ProductSearchRequest) gorm.DB {
	if criteria.Name != nil && *criteria.Name != "" {
		db = *db.Where("name like ?", "%"+*criteria.Name+"%")
	}
	if criteria.Remark != nil && *criteria.Remark != "" {
		db = *db.Where("remark like ?", "%"+*criteria.Remark+"%")
	}
	return db
}

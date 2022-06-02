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

type CustomerService struct {
	Container container.Container
}

func NewCustomerService(container container.Container) *CustomerService {
	sv := new(CustomerService)
	sv.Container = container
	return sv
}

func (sv CustomerService) Search(criteria model.CustomerSearchRequest) (*model.SearchResult[model.CustomerDataReponse], error) {
	repo, err := core.Resolve[repository.CustomerRepository](sv.Container)
	if err != nil {
		return nil, err
	}

	result := model.SearchResult[model.CustomerDataReponse]{}
	query := sv.FilterCustomer(*repo.GetAll(), criteria)
	var itemCount int64
	core.HandleDBError(query.Count(&itemCount))
	result.ItemCount = int(itemCount)

	query = core.Sort(query, criteria.SortBy, criteria.SortDirection)
	query = core.Page(query, criteria.PageLength, criteria.Page)
	resultData, err := core.MapQuery(query, func(x entity.Customer) model.CustomerDataReponse {
		return model.CustomerDataReponse{
			ID:        x.ID,
			FirstName: x.FirstName,
			LastName:  x.LastName,
		}
	})
	if err != nil {
		return nil, err
	}
	result.ResultData = resultData

	return &result, nil
}

func (sv CustomerService) AddCustomer(data model.CustomerUpdateRequest) (*uuid.UUID, error) {
	repo, err := core.Resolve[repository.CustomerRepository](sv.Container)
	if err != nil {
		return nil, err
	}
	customer := new(entity.Customer)
	customer.ID = uuid.New()
	customer.FirstName = data.FirstName
	customer.LastName = data.LastName
	customer.CreatedBy = core.Ptr("Testt")
	customer.CreatedDate = core.Ptr(time.Now())
	repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Save(*customer)
		return nil
	})
	return &customer.ID, nil
}

func (sv CustomerService) UpdateCustomer(id uuid.UUID, data model.CustomerUpdateRequest) error {
	repo, err := core.Resolve[repository.CustomerRepository](sv.Container)
	if err != nil {
		return err
	}
	customer, err := repo.Get(id)
	if err != nil {
		return err
	}
	customer.FirstName = data.FirstName
	customer.LastName = data.LastName
	customer.UpdatedBy = core.Ptr("Testt")
	customer.UpdatedDate = core.Ptr(time.Now())
	err = repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Save(*customer)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (sv CustomerService) RemoveCustomer(id uuid.UUID) error {
	repo, err := core.Resolve[repository.CustomerRepository](sv.Container)
	if err != nil {
		return err
	}
	customer, err := repo.Get(id)
	if err != nil {
		return err
	}
	err = repo.DB.Transaction(func(tx *gorm.DB) error {
		repo.Remove(*customer)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (sv CustomerService) FilterCustomer(db gorm.DB, criteria model.CustomerSearchRequest) gorm.DB {
	if criteria.Name != nil && *criteria.Name != "" {
		db = *db.Where("first_name like ? or last_name like ?", "%"+*criteria.Name+"%", "%"+*criteria.Name+"%")
	}
	return db
}

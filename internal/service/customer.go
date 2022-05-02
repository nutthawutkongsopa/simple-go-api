package service

import (
	"log"
	"test-api/internal/core"
	"test-api/internal/data/entity"
	"test-api/internal/data/repository"
	"test-api/internal/model"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAll() *gorm.DB
}

type CustomerService struct {
	Container container.Container
}

func (sv CustomerService) Search(criteria model.CustomerSearchRequest) model.SearchResult[model.CustomerDataReponse] {
	var repo repository.CustomerRepository
	err := sv.Container.Resolve(&repo)
	if err != nil {
		log.Fatal(err)
	}
	result := model.SearchResult[model.CustomerDataReponse]{}
	query := FilterCustomer(*repo.GetAll(), criteria)
	query.Count(&result.ItemCount)

	query = core.Sort(query, criteria.SortBy, criteria.SortDirection)
	resultData := core.MapQuery(query, func(x entity.Customer) model.CustomerDataReponse {
		return model.CustomerDataReponse{
			ID:        x.ID,
			FirstName: x.FirstName,
			LastName:  x.LastName,
		}
	})
	result.ResultData = resultData

	return result
}

func FilterCustomer(db gorm.DB, criteria model.CustomerSearchRequest) gorm.DB {
	if criteria.Name != nil && *criteria.Name != "" {
		db = *db.Where("first_name like ? or last_name like ?", *criteria.Name, *criteria.Name)
	}
	return db
}

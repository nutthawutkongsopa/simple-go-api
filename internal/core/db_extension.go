package core

import (
	"fmt"

	"gorm.io/gorm"
)

func Page(db gorm.DB, pageLength *int, page *int) gorm.DB {
	if pageLength != nil && page != nil {
		_page := IIF(*page > 0, *page, 1)
		_length := IIF(*pageLength > 0, *pageLength, 1)
		offset := (_page - 1) * _length

		return *db.Offset(offset).Limit(_length)
	}
	return db
}

func Sort(db gorm.DB, sortBy *string, sortDirection *string) gorm.DB {
	if sortBy != nil && sortDirection != nil {
		_direction := IIF(*sortDirection == "desc", "desc", "asc")

		return *db.Order(fmt.Sprintf("%s %s", *sortBy, _direction))
	}
	return db
}

func HandleDBError(db *gorm.DB) (*gorm.DB, error) {
	if db.Error != nil {
		return db, db.Error
	}

	return db, nil
}

func MapQuery[TQuery interface{}, TResult interface{}](db gorm.DB, selector Selector[TQuery, TResult]) ([]TResult, error) {
	var result []TResult = []TResult{}
	var items []TQuery

	_, err := HandleDBError(db.Find(&items))
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		result = append(result, selector(item))
	}
	return result, nil
}

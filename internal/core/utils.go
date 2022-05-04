package core

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type Selector[T interface{}, TResult interface{}] func(t T) TResult

func CheckFileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func IIF[T interface{}](condition bool, consequent T, alternative T) T {
	if condition {
		return consequent
	} else {
		return alternative
	}
}

func GetCurrentDirectory() string {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)
	exPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return exPath
}

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

func MapQuery[TQuery interface{}, TResult interface{}](db gorm.DB, selector Selector[TQuery, TResult]) []TResult {
	var result []TResult
	var items []TQuery

	db.Find(&items)
	for _, item := range items {
		result = append(result, selector(item))
	}
	return result
}

func Resolve[TResult interface{}](c container.Container) TResult {
	var result TResult
	err := c.Resolve(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func Ptr[T interface{}](value T) *T {
	return &value
}

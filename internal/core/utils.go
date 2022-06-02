package core

import (
	"errors"
	"io/fs"
	"os"

	"github.com/golobby/container/v3"
)

type Selector[TSource interface{}, TResult interface{}] func(t TSource) TResult

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

func GetCurrentDirectory() (string, error) {
	return os.Getwd()
}

func Resolve[TResult interface{}](c container.Container) (*TResult, error) {
	var result TResult
	err := c.Resolve(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Ptr[T interface{}](value T) *T {
	return &value
}

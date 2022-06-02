package app

import (
	"test-api/internal/data/repository"
	"test-api/internal/service"

	"github.com/golobby/container/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupContainer() (container.Container, error) {
	container := container.New()

	err := RegistConfig(container)
	if err != nil {
		return nil, err
	}

	err = RegistDBConnection(container)

	if err != nil {
		return nil, err
	}

	err = RegistRepository(container)

	if err != nil {
		return nil, err
	}

	err = RegistCustomService(container)

	if err != nil {
		return nil, err
	}

	return container, nil
}

func RegistConfig(container container.Container) error {
	env, err := GetEnv()
	if err != nil {
		return err
	}

	container.Singleton(func() Environment {
		return *env
	})

	appsettings, err := GetAppSettigs()
	if err != nil {
		return err
	}

	container.Singleton(func() AppSettings {
		return *appsettings
	})

	return nil
}

func RegistDBConnection(container container.Container) error {
	var appsettings AppSettings
	container.Resolve(&appsettings)

	return container.Singleton(func() (*gorm.DB, error) {
		conn := appsettings.DBConnectionString
		return gorm.Open(postgres.Open(conn))
	})
}

func RegistRepository(container container.Container) error {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return err
	}

	if err = container.Transient(func() repository.CustomerRepository {
		return *repository.NewCustomerRepository(*db)
	}); err != nil {
		return err
	}

	if err = container.Transient(func() repository.ProductRepository {
		return *repository.NewProductRepository(*db)
	}); err != nil {
		return err
	}

	return nil
}

func RegistCustomService(container container.Container) error {
	if err := container.Transient(func() service.CustomerService {
		return service.CustomerService{
			Container: container,
		}
	}); err != nil {
		return err
	}

	if err := container.Transient(func() service.ProductService {
		return service.ProductService{
			Container: container,
		}
	}); err != nil {
		return err
	}

	return nil
}

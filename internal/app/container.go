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

	RegistDBConnection(container)

	RegistRepository(container)

	RegistCustomService(container)

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

func RegistDBConnection(container container.Container) {
	var appsettings AppSettings
	container.Resolve(&appsettings)

	container.Singleton(func() *gorm.DB {
		conn := appsettings.DBConnectionString
		db, err := gorm.Open(postgres.Open(conn))

		if err != nil {
			panic("failed to connect database")
		}

		return db
	})
}

func RegistRepository(container container.Container) {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		panic(err)
	}
	container.Transient(func() repository.CustomerRepository {
		return *repository.NewCustomerRepository(*db)
	})
	container.Transient(func() repository.ProductRepository {
		return *repository.NewProductRepository(*db)
	})
}

func RegistCustomService(container container.Container) {
	container.Transient(func() service.CustomerService {
		return service.CustomerService{
			Container: container,
		}
	})
	container.Transient(func() service.ProductService {
		return service.ProductService{
			Container: container,
		}
	})
}

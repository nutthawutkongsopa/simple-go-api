package main

import (
	"fmt"
	"log"
	"test-api/internal/app"
	"test-api/internal/model"
	"test-api/internal/service"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p@ssw0rd"
	dbname   = "test_api_db"
)

func main() {
	container, err := app.SetupContainer() //app.SetupContainer()
	if err != nil {
		log.Fatal(err)
	}
	var xx service.CustomerService
	container.Resolve(&xx)
	yy := xx.Search(model.CustomerSearchRequest{})
	fmt.Printf(yy.ResultData[0].ID.String())
	_ = yy
}

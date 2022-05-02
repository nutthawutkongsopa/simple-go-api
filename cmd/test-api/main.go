package main

import (
	"fmt"
	"test-api/internal/app"
	"test-api/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	defer deferring()
	container, err := app.SetupContainer()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.InitRoute(r, container)
	r.Run()
}

func deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}

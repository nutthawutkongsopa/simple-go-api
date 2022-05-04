package router

import (
	"test-api/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

func SetCustomerRoutes(g gin.IRouter, container container.Container) {
	var controller = new(controller.CustomerController)
	controller.Container = container

	r := g.Group("/customers")
	r.GET("/", controller.SearchCustomer)
	r.POST("/", controller.AddCustomer)
	r.PUT("/:id", controller.UpdateCustomer)
	r.DELETE("/:id", controller.RemoveCustomer)
}

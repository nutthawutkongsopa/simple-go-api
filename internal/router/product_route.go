package router

import (
	"test-api/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

func SetProductRoutes(g gin.IRouter, container container.Container) {
	var controller = new(controller.ProductController)
	controller.Container = container

	r := g.Group("/products")
	r.GET("/", controller.SearchProduct)
	r.POST("/", controller.AddProduct)
	r.PUT("/:id", controller.UpdateProduct)
	r.DELETE("/:id", controller.RemoveProduct)
}

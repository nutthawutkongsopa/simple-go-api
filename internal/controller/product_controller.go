package controller

import (
	"net/http"
	"test-api/internal/model"
	"test-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"github.com/google/uuid"
)

type ProductController struct {
	Container container.Container
}

func NewProductController(c container.Container) *ProductController {
	return &ProductController{
		Container: c,
	}
}

// @Tags         product
// @Summary      SearchProduct
// @Accept       json
// @Produce      json
// @Param        criteria    query     model.ProductSearchRequest  false  "criteria for retrive"
// @Success      200  {object}   model.APIResult
// @Router       /products [get]
func (controller ProductController) SearchProduct(c *gin.Context) {
	var service service.ProductService
	controller.Container.Resolve(&service)
	var criteria model.ProductSearchRequest
	if err := c.BindQuery(&criteria); err != nil {
		panic(err)
	}
	result := service.Search(criteria)
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(result))
}

// @Tags         product
// @Summary      AddProduct
// @Accept       json
// @Produce      json
// @Param        data    body     model.ProductUpdateRequest  false  "data to update"
// @Success      200  {object}   model.APIResult
// @Router       /products [post]
func (controller ProductController) AddProduct(c *gin.Context) {
	var service service.ProductService
	controller.Container.Resolve(&service)
	var data model.ProductUpdateRequest
	c.BindJSON(&data)
	id := service.AddProduct(data)
	result := struct {
		Id uuid.UUID `json:"id"`
	}{Id: id}
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(result))
}

// @Tags         product
// @Summary      UpdateProduct
// @Accept       json
// @Produce      json
// @Param		 id		 path 	  string	true "Product ID" Format(uuid)
// @Param        data    body     model.ProductUpdateRequest  false  "data to update"
// @Success      200  {object}   model.APIResult
// @Router       /products/{id} [put]
func (controller ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var service service.ProductService
	controller.Container.Resolve(&service)
	var data model.ProductUpdateRequest
	c.BindJSON(&data)
	service.UpdateProduct(uuid.MustParse(id), data)
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(nil))
}

// @Tags         product
// @Summary      RemoveProduct
// @Accept       json
// @Produce      json
// @Param		 id		 path 	  string	true "Product ID" Format(uuid)
// @Success      200  {object}   model.APIResult
// @Router       /products/{id} [delete]
func (controller ProductController) RemoveProduct(c *gin.Context) {
	id := c.Param("id")
	var service service.ProductService
	controller.Container.Resolve(&service)
	service.RemoveProduct(uuid.MustParse(id))
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(nil))
}

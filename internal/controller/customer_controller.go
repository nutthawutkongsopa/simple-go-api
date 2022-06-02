package controller

import (
	"net/http"
	"test-api/internal/core"
	"test-api/internal/model"
	"test-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"github.com/google/uuid"
)

type CustomerController struct {
	Container container.Container
}

func NewCustomerController(c container.Container) *CustomerController {
	return &CustomerController{
		Container: c,
	}
}

// @Tags         customer
// @Summary      SearchCustomer
// @Accept       json
// @Produce      json
// @Param        criteria    query     model.CustomerSearchRequest  false  "criteria for retrive"
// @Success      200  {object}   model.APIResult
// @Router       /customers [get]
func (controller CustomerController) SearchCustomer(c *gin.Context) {
	var service service.CustomerService
	controller.Container.Resolve(&service)
	var criteria model.CustomerSearchRequest
	if err := c.BindQuery(&criteria); err != nil {
		core.HandleAPIError(c, err)
		return
	}
	result, err := service.Search(criteria)
	if err != nil {
		core.HandleAPIError(c, err)
		return
	}
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(result))
}

// @Tags         customer
// @Summary      AddCustomer
// @Accept       json
// @Produce      json
// @Param        data    body     model.CustomerUpdateRequest  false  "data to update"
// @Success      200  {object}   model.APIResult
// @Router       /customers [post]
func (controller CustomerController) AddCustomer(c *gin.Context) {
	var service service.CustomerService
	controller.Container.Resolve(&service)
	var data model.CustomerUpdateRequest
	if err := c.BindJSON(&data); err != nil {
		core.HandleAPIError(c, err)
		return
	}
	id, err := service.AddCustomer(data)
	if err != nil {
		core.HandleAPIError(c, err)
		return
	}
	result := struct {
		Id uuid.UUID `json:"id"`
	}{Id: *id}
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(result))
}

// @Tags         customer
// @Summary      UpdateCustomer
// @Accept       json
// @Produce      json
// @Param		 id		 path 	  string	true "Customer ID" Format(uuid)
// @Param        data    body     model.CustomerUpdateRequest  false  "data to update"
// @Success      200  {object}   model.APIResult
// @Router       /customers/{id} [put]
func (controller CustomerController) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var service service.CustomerService
	controller.Container.Resolve(&service)
	var data model.CustomerUpdateRequest
	if err := c.BindJSON(&data); err != nil {
		core.HandleAPIError(c, err)
		return
	}
	service.UpdateCustomer(uuid.MustParse(id), data)
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(nil))
}

// @Tags         customer
// @Summary      RemoveCustomer
// @Accept       json
// @Produce      json
// @Param		 id		 path 	  string	true "Customer ID" Format(uuid)
// @Success      200  {object}   model.APIResult
// @Router       /customers/{id} [delete]
func (controller CustomerController) RemoveCustomer(c *gin.Context) {
	id := c.Param("id")
	var service service.CustomerService
	controller.Container.Resolve(&service)
	service.RemoveCustomer(uuid.MustParse(id))
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(nil))
}

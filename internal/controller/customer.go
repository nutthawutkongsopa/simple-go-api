package controller

import (
	"net/http"
	"test-api/internal/model"
	"test-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

type CustomerController struct {
	Container container.Container
}

func NewCustomerController(c container.Container) *CustomerController {
	return &CustomerController{
		Container: c,
	}
}

// @Tags Customer
// @Router /customers [get]
// @Accept json
// @Param id  path int true "Customer ID"
func (controller CustomerController) SearchCustomer(c *gin.Context) {
	var service service.CustomerService
	controller.Container.Resolve(&service)
	result := service.Search(model.CustomerSearchRequest{})
	c.JSON(http.StatusOK, model.NewAPIResultSuccess(result))
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

func InitRoute(e *gin.Engine, container container.Container) {
	SetCustomerRoutes(e, container)
	SetProductRoutes(e, container)
}

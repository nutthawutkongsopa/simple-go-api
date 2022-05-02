package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

func InitRoute(r gin.IRouter, container container.Container) {
	SetCutomerRoutes(r, container)
}

package main

import (
	"strings"
	"test-api/internal/app"
	"test-api/internal/core"
	"test-api/internal/middleware"
	"test-api/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//defer deferring()
	container, err := app.SetupContainer()
	if err != nil {
		panic(err)
	}

	var settings app.AppSettings
	container.Resolve(&settings)

	if settings.GINMode != "" {
		gin.SetMode(settings.GINMode)
	}

	e := gin.New()

	app.InitSwagger(e, settings.AppHost)

	e.Use(middleware.ContextAccessMiddleware)
	core.InitSessionData(func() *gin.Context {
		ctx := middleware.GetCurrentContxt()
		return ctx
	})

	e.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		app.HandleAPIError(c, err)
		c.Next()
	}))

	router.InitRoute(e, container)

	appAddrs := strings.Split(settings.ApplicateionURL, ";")
	e.Run(appAddrs...)
}

// func deferring() {
// 	if err := recover(); err != nil {
// 		fmt.Println("An error occurred:", err)
// 	}
// }

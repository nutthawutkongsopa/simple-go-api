package main

import (
	"strings"
	"test-api/internal/app"
	"test-api/internal/core"
	"test-api/internal/middleware"
	"test-api/internal/router"

	"github.com/gin-gonic/gin"
)

// @title API document title
// @version version(1.0)
// @description Description of specifications
// @Precautions when using termsOfService specifications

// @contact.name API supporter
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name license(Mandatory)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

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

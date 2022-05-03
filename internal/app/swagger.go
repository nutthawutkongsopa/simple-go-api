package app

import (
	"test-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwagger(e *gin.Engine, baseURL string) {
	docs.SwaggerInfo.Title = "Test API"
	docs.SwaggerInfo.Description = "This is a example golang restful api."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = baseURL
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

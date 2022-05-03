package middleware

import (
	"github.com/gin-gonic/gin"
)

var currentContext *gin.Context

var ContextAccessMiddleware gin.HandlerFunc = func(ctx *gin.Context) {
	currentContext = ctx
	ctx.Next()
}

func GetCurrentContxt() *gin.Context {
	return currentContext
}

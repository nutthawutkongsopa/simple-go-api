package app

import (
	"test-api/internal/core"
	"test-api/internal/model"

	"github.com/gin-gonic/gin"
)

func HandleAPIError(c *gin.Context, err interface{}) {
	switch err.(type) {
	case *core.DataErrorException:
		{
			e := err.(*core.DataErrorException)
			c.JSON(400, model.NewAPIResultError(400, *e.ErrorCode, *e.Message, *e.Details))
		}
	case core.DataErrorException:
		{
			e := err.(core.DataErrorException)
			c.JSON(400, model.NewAPIResultError(400, *e.ErrorCode, *e.Message, *e.Details))
		}
	default:
		{
			c.JSON(500, model.NewAPIResultError(500, "12345", "An error occurred", ""))
		}
	}
}

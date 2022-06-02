package core

import (
	"fmt"
	"test-api/internal/model"

	"github.com/gin-gonic/gin"
)

func HandleAPIError(c *gin.Context, err interface{}) {
	switch err.(type) {
	case *DataErrorException:
		{
			e := err.(*DataErrorException)
			c.JSON(400, model.NewAPIResultError(400, *e.ErrorCode, *e.Message, *e.Details))
		}
	case DataErrorException:
		{
			e := err.(DataErrorException)
			c.JSON(400, model.NewAPIResultError(400, *e.ErrorCode, *e.Message, *e.Details))
		}
	default:
		{
			c.JSON(500, model.NewAPIResultError(500, "12345", "An error occurred", fmt.Sprintln(err)))
		}
	}
}

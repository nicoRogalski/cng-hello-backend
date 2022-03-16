package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/pkg/errors"
)

func Handle(c *gin.Context, err error) {
	switch err := err.(type) {
	case errors.ErrNotFound:
		c.IndentedJSON(err.Code, err.Message)
		return
	case errors.ErrInternalServer:
		c.IndentedJSON(err.Code, err.Message)
		return
	default:
		c.IndentedJSON(500, "Internal server error")
		return
	}
}

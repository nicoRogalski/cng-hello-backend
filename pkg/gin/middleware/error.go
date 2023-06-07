package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/pkg/errors"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
	e := ctx.Errors
	if e != nil {
		handle(ctx, ctx.Errors.Last())
	}
}

func handle(c *gin.Context, err error) {
	switch err := err.(type) {
	case errors.ErrNotFound:
		c.AbortWithStatusJSON(err.Code, err)
		return
	case errors.ErrInternalServer:
		c.AbortWithStatusJSON(err.Code, err)
		return
	default:
		e := errors.ErrInternalServer{Code: 500, Message: "Internal server error"}
		c.AbortWithStatusJSON(e.Code, e)
		return
	}
}

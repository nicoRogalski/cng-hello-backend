package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rogalni/cng-hello-backend/internal/service"
)

const BEARER_SCHEMA = "Bearer"
const AUTH_HEADER = "Authorization"

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AUTH_HEADER)
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			//TODO set claims in gin context
			// c.Set()
			fmt.Println(claims)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

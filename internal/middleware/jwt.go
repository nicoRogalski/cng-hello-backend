package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rogalni/cng-hello-backend/internal/service"
)

const BEARER_SCHEMA = "Bearer "
const AUTH_HEADER = "Authorization"

func ValidateJWT(c *gin.Context) {
	authHeader := c.GetHeader(AUTH_HEADER)
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := service.ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		// Adding groups to context
		g := claims["groups"].([]interface{})
		c.Set("groups", g)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

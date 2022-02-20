package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
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
	token, err := auth.ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	//TODO: Outsource in pkg
	// Adding groups to context
	g := claims["groups"].([]interface{})
	c.Set("groups", g)

}

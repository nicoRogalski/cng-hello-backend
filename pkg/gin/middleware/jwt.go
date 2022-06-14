package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
	gauth "github.com/rogalni/cng-hello-backend/pkg/gin/auth"
)

const (
	authHeader = "Authorization"
)

func ValidateJWT(c *gin.Context) {
	ts, err := auth.ExtractJWT(c.GetHeader(authHeader))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	jwks := auth.Jwks
	if jwks == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	token, err := jwt.Parse(ts, jwks.Keyfunc)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Set(gauth.ClaimsFieldName, claims)
}

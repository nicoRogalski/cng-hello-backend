package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
	gauth "github.com/rogalni/cng-hello-backend/pkg/gin/auth"
)

func ValidateJWT(c *gin.Context) {
	tokenString, err := auth.ExtractJWT(c.GetHeader(auth.AuthHeader))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}
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
	c.Set(gauth.ClaimsFieldName, claims)
}

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	ClaimsFieldName = "claims"
)

// Gets the claims from the context after the middleware placed it inside on setup
func GetClaims(c *gin.Context) (claims jwt.MapClaims, exists bool) {
	cl, e := c.Get(ClaimsFieldName)
	if !e {
		return nil, false
	}
	return cl.(jwt.MapClaims), true
}

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	ClaimsFieldName = "claims"
)

func GetJWTRoles(c *gin.Context) []string {
	claims, e := getClaims(c)
	if !e {
		return []string{}
	}
	gs, e := claims["groups"].([]interface{})
	if !e {
		return []string{}
	}

	g := make([]string, len(gs))
	for i, v := range gs {
		g[i] = v.(string)
	}
	return g
}

// Gets the claims from the context after the middleware placed it inside on setup
func getClaims(c *gin.Context) (claims jwt.MapClaims, exists bool) {
	cl, e := c.Get(ClaimsFieldName)
	if !e {
		return nil, false
	}
	return cl.(jwt.MapClaims), true
}

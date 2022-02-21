package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/pkg/gin/auth"
)

func GetJWTRoles(c *gin.Context) []string {
	claims, e := auth.GetClaims(c)
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

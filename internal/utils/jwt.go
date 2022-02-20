package utils

import (
	"github.com/gin-gonic/gin"
)

func GetJWTRoles(c *gin.Context) []string {
	gs, e := c.Get("groups")
	if !e {
		return []string{}
	}
	gg := gs.([]interface{})
	g := make([]string, len(gg))
	for i, v := range gg {
		g[i] = v.(string)
	}
	return g
}

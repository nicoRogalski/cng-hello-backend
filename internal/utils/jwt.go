package utils

import "github.com/gin-gonic/gin"

func GetJWTRole(c *gin.Context) string {
	r, e := c.Get("ROLE")
	if !e {
		return ""
	}
	return r.(string)
}

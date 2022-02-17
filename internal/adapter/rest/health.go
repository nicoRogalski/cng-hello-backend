package rest

import "github.com/gin-gonic/gin"

type Health struct {
	Status string `json:"status"`
}

func HealthHandler(c *gin.Context) {
	
	c.JSON(200, &Health{
		Status: "UP",
	})
}

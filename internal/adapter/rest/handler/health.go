package handler

import "github.com/gin-gonic/gin"

type Health struct {
	Status string `json:"status"`
}

func GetHealth(c *gin.Context) {
	c.JSON(200, &Health{
		Status: "UP",
	})
}

func GetReadiness(c *gin.Context) {
	c.JSON(200, &Health{
		Status: "UP",
	})
}

func GetLiveness(c *gin.Context) {
	c.JSON(200, &Health{
		Status: "UP",
	})
}

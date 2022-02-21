package health

import "github.com/gin-gonic/gin"

type Health struct {
	Status string `json:"status"`
}

func For(r *gin.Engine) {
	health := r.Group("/health")
	{
		health.GET("/", getHealth)
		health.GET("/readiness", getReadiness)
		health.GET("/liveness", getLiveness)
	}
}

func getHealth(c *gin.Context) {
	c.JSON(200, &Health{
		Status: "UP",
	})
}

func getReadiness(c *gin.Context) {
	c.JSON(200, &Health{
		Status: "UP",
	})
}

func getLiveness(c *gin.Context) {
	c.JSON(200, &Health{
		Status: "UP",
	})
}

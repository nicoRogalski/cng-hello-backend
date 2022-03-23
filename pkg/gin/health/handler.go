package health

import "github.com/gin-gonic/gin"

type Health struct {
	Status     string            `json:"status"`
	Components map[string]string `json:"components,omitempty"`
}
type Status func() Health

func For(r *gin.Engine, extendedStatus Status) *gin.Engine {
	health := r.Group("/health")
	{
		health.GET("/", getHealth)
		health.GET("/readiness", func(c *gin.Context) {
			c.IndentedJSON(200, extendedStatus())
		})
		health.GET("/liveness", func(c *gin.Context) {
			c.IndentedJSON(200, extendedStatus())
		})
	}
	return r
}

func getHealth(c *gin.Context) {
	c.JSON(200, &Health{
		Status: "UP",
	})
}

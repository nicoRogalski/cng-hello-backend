package health

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	rg *gin.RouterGroup
}
type StatusFunc func() Health

func For(r *gin.Engine) *HealthHandler {
	rg := r.Group("/health")
	{
		rg.GET("", getHealth)
	}
	hh := &HealthHandler{
		rg: rg,
	}
	return hh
}

func (h *HealthHandler) WithReadiness(sf StatusFunc) *HealthHandler {
	h.rg.GET("/readiness", func(c *gin.Context) {
		es := sf()
		for _, v := range es.Components {
			if v.Status == DOWN {
				es.Status = DOWN
				break
			}
		}
		c.IndentedJSON(es.Code, es)
	})
	return h
}

func (h *HealthHandler) WithLiveness(sf StatusFunc) *HealthHandler {
	h.rg.GET("/liveness", func(c *gin.Context) {
		es := sf()
		es.Status = UP
		for _, v := range es.Components {
			if v.Status == DOWN {
				es.Status = DOWN
				break
			}
		}
		c.IndentedJSON(es.Code, es)
	})
	return h
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(200, &Health{
		Status: UP,
	})
}

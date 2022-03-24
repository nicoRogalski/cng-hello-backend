package health

import (
	"github.com/gin-gonic/gin"
)

var (
	UP   = "UP"
	DOWN = "DOWN"
)

type HealthHandler struct {
	rg *gin.RouterGroup
}

type StatusFunc func() Health

type Health struct {
	Status     string      `json:"status"`
	Code       int         `json:"-"`
	Components []Component `json:"components,omitempty"`
}

type Component struct {
	Name   string
	Status string
}

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

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/utils"
	"github.com/rogalni/cng-hello-backend/pkg/gin/tracer"
	zlog "github.com/rogalni/cng-hello-backend/pkg/log"
	"github.com/rs/zerolog/log"
)

type HelloMessage struct {
	Text string `json:"text"`
}

func GetHello(c *gin.Context) {
	span := tracer.Start(c.Request.Context(), "handler.GetHello")
	defer span.End()
	//Example with span in context
	zlog.Ctx(c.Request.Context()).Info().Msg("Get Hello with trace infos from context")
	//Example with span
	zlog.Span(span).Info().Msg("Get Hello with trace infos from span")

	c.IndentedJSON(200, &HelloMessage{
		Text: "Welcome to cloud native go",
	})
}

func GetHelloSecure(c *gin.Context) {
	roles := utils.GetJWTRoles(c)
	log.Info().Msgf("Authorized with role: %s", roles)
	span := tracer.Start(c.Request.Context(), "GetHelloSecure")
	defer span.End()
	c.IndentedJSON(200, &HelloMessage{
		Text: "Welcome to cloud native go with jwt auth",
	})
}

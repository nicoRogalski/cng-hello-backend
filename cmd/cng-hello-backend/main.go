package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/v1/handler"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
	"github.com/rogalni/cng-hello-backend/pkg/gin/health"
	"github.com/rogalni/cng-hello-backend/pkg/gin/middleware"
	"github.com/rogalni/cng-hello-backend/pkg/otel"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()

	shutdown := otel.Setup(ctx, cfg)
	defer shutdown(ctx)

	go auth.Setup(cfg.JwkSetUri)
	gdb := postgres.InitConnection(cfg.PostgresHost, cfg.PostgresUser, cfg.PostresPassword, cfg.PostgresDb, cfg.PostgresPort)

	server := NewServer(cfg, gdb)
	server.Run()
}

type Server struct {
	cfg *config.Config
	r   *gin.Engine
	gdb *gorm.DB
}

func NewServer(cfg *config.Config, gdb *gorm.DB) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	s := &Server{
		cfg: cfg,
		r:   r,
		gdb: gdb,
	}
	s.setupRoutes()
	return s
}

func (s *Server) Run() {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.Port),
		Handler: s.r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			otelzap.S().Fatalf("Error serve http server %s\n", err)
		}
	}()
	otelzap.S().Info("Server Started on port: ", s.cfg.Port)

	<-done
	otelzap.L().Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		otelzap.S().Fatalf("Server Shutdown Failed:%+v", err)
	}
	otelzap.L().Info("Server Exited Properly")
}

func (s *Server) setupRoutes() {
	s.r.Use(gin.Recovery())
	s.r.Use(middleware.ErrorHandler)

	sf := func() (h health.Health) {
		return serverStatus(s.gdb)
	}
	health.For(s.r).
		WithLiveness(sf).
		WithReadiness(sf)

	api := s.r.Group("/api")
	// Tracing and endpoint logging is attached to "/api" route
	api.Use(otelgin.Middleware(s.cfg.ServiceName))
	api.Use(middleware.Logger(s.cfg.ServiceName))
	api.Use(gzip.Gzip(gzip.DefaultCompression))

	// "v1" simulates a real world example of endpoints
	v1 := api.Group("/v1")
	{
		m := v1.Group("/messages")
		{
			mh := handler.NewMessageHandler(s.gdb)
			m.GET("", mh.GetMessages)
			m.GET("/:id", mh.GetMessageById)
			m.POST("", mh.CreateMessage)
			m.DELETE("/:id", middleware.ValidateJWT, mh.DeleteMessage)
		}
	}
}

func serverStatus(gdb *gorm.DB) (h health.Health) {
	phc := health.Component{Name: "postgres"}
	if db, err := gdb.DB(); err != nil {
		phc.Status = health.DOWN
	} else if err := db.Ping(); err != nil {
		phc.Status = health.DOWN
	} else {
		phc.Status = health.UP
	}

	h.Components = []health.Component{phc}
	return
}

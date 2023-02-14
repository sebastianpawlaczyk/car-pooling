package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/car-pooling/internal/cfg"
	"github.com/car-pooling/internal/endpoint/healthCheck_api"
	"github.com/car-pooling/internal/endpoint/pooling_api"
)

type HttpServer struct {
	echoServer         *echo.Echo
	config             *cfg.Config
	healthCheckHandler *healthCheck_api.HealthCheckHandler
	poolingHandler     *pooling_api.PoolingHandler
}

func NewHttpServer(config *cfg.Config, healthCheckHandler *healthCheck_api.HealthCheckHandler, poolingHandler *pooling_api.PoolingHandler) *HttpServer {
	return &HttpServer{
		config:             config,
		healthCheckHandler: healthCheckHandler,
		poolingHandler:     poolingHandler,
	}
}

func (s *HttpServer) init() error {
	s.echoServer = echo.New()
	s.echoServer.HidePort = true
	s.echoServer.HideBanner = true
	s.echoServer.Use(middleware.CORS())

	s.register()
	return nil
}

func (s *HttpServer) Start() error {
	err := s.init()
	if err != nil {
		return err
	}

	log.Info().Msgf("Starting server on port %v", s.config.Port)
	return s.echoServer.Start(fmt.Sprintf(":%v", s.config.Port))
}

func (s *HttpServer) register() {
	s.healthCheckHandler.Register(s.echoServer)
	s.poolingHandler.Register(s.echoServer)
}

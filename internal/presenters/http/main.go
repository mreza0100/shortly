package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/ports/services"
)

func NewHttpServer(port string, isDev bool, services *services.Services) *server {
	return &server{port: port, isDev: isDev, service: services}
}

type server struct {
	port      string
	isDev     bool
	ginClient *gin.Engine
	service   *services.Services
}

func (s *server) ListenAndServe() <-chan error {
	s.ginClient = gin.Default()
	s.registerRoutes()

	if s.isDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	errCh := make(chan error)
	go func(errCh chan error) {
		addr := fmt.Sprint(":", s.port)
		if err := s.ginClient.Run(addr); err != nil {
			errCh <- err
		}
	}(errCh)

	return errCh
}

func (s *server) registerRoutes() {
	registerUserRoutes(s.ginClient, s.service.User)
}

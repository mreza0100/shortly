package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func NewHttpServer(port string, service interface{}) *server {
	if port == "" {
		port = defaultPort
	}
	return &server{port: port, service: service}
}

type server struct {
	port      string
	ginClient *gin.Engine
	service   interface{}
}

func (s *server) Serve() {
	s.ginClient = gin.Default()
	s.ginClient.Use(gin.Recovery())
	s.ginClient.Use(gin.Logger())
	s.ginClient.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: gin.DefaultErrorWriter,
	}))
	s.registerRoutes()

	go func() {
		port := fmt.Sprint(":", s.port)
		if err := s.ginClient.Run(port); err != nil {
			panic(err)
		}
	}()
}

func (s *server) registerRoutes() {
}

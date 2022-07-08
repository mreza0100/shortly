package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mreza0100/shortly/internal/ports/services"
	"github.com/mreza0100/shortly/pkg/jwt"
)

type NewHttpServerOpts struct {
	Port     string
	IsDev    bool
	JwtUtils jwt.JWTHelper
	Services *services.Services
}

func NewHttpServer(opts NewHttpServerOpts) *server {
	return &server{
		port:     opts.Port,
		isDev:    opts.IsDev,
		jwtUtils: opts.JwtUtils,
		services: opts.Services,
	}
}

type server struct {
	port      string
	isDev     bool
	ginClient *gin.Engine
	services  *services.Services
	jwtUtils  jwt.JWTHelper
}

func (s *server) ListenAndServe() <-chan error {
	errCh := make(chan error)
	s.ginClient = gin.Default()
	s.registerRoutes()

	if err := s.ginClient.SetTrustedProxies(nil); err != nil {
		go func(errCh chan error) { errCh <- err }(errCh)
		goto exit
	}

	if s.isDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	go func(errCh chan error) {
		addr := fmt.Sprint(":", s.port)
		errCh <- s.ginClient.Run(addr)
	}(errCh)

exit:
	return errCh
}

func (s *server) registerRoutes() {
	registerUserRoutes(s.ginClient, s.services.User)
	registerLinkRoutes(s.ginClient, s.jwtUtils, s.services.Link)
}

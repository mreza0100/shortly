package services

import (
	"github.com/mreza0100/shortly/internal/ports/services"
)

func NewAuthenticationService() services.AuthenticationServicePort {
	return &authentication{}
}

type authentication struct{}

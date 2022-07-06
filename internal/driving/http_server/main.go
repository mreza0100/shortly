package httpserver

import (
	"github.com/mreza0100/shortly/internal/ports/driving"
)

func New() driving.HTTPServerPort {
	return &HTTPServer{}
}

type HTTPServer struct{}

package services

import (
	"context"

	"github.com/mreza0100/shortly/internal/ports"
)

// Health Service Dependencies
type HealthServiceDep struct {
	StorageRead ports.StorageReadPort
}

// Constructor of health service
func NewHealthService(opt *HealthServiceDep) ports.HealthServicePort {
	return &health{
		storageRead: opt.StorageRead,
	}
}

// Health service implementation
type health struct {
	storageRead ports.StorageReadPort
}

// check storage health
func (h *health) CheckHealth(ctx context.Context) bool {
	// storage usually takes a long time to start
	return h.storageRead.HealthCheck(ctx)
}

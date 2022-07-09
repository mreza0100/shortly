package services

import (
	"context"

	"github.com/mreza0100/shortly/internal/ports"
)

//
type HealthServiceDep struct {
	StorageRead ports.StorageReadPort
}

func NewHealthService(opt *HealthServiceDep) ports.HealthServicePort {
	return &health{
		storageRead: opt.StorageRead,
	}
}

type health struct {
	storageRead ports.StorageReadPort
}

func (h *health) CheckHealth(ctx context.Context) bool {
	// check storage health
	// storage usually takes a long time to start
	return h.storageRead.HealthCheck(ctx)
}

package services

import (
	"context"

	"github.com/mreza0100/shortly/internal/ports"
)

type HealthServiceOptions struct {
	CassandraRead ports.StorageReadPort
}

func NewHealthService(opt *HealthServiceOptions) ports.HealthServicePort {
	return &health{
		CassandraRead: opt.CassandraRead,
	}
}

type health struct {
	CassandraRead ports.StorageReadPort
}

func (h *health) CheckHealth(ctx context.Context) bool {
	return h.CassandraRead.HealthCheck(ctx)
}

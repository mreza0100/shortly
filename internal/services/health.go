package services

import (
	"context"

	"github.com/mreza0100/shortly/internal/ports"
)

func NewHealthService(cassandraRead ports.CassandraReadPort) ports.HealthServicePort {
	return &health{
		CassandraRead: cassandraRead,
	}
}

type health struct {
	CassandraRead ports.CassandraReadPort
}

func (h *health) CheckHealth(ctx context.Context) bool {
	return h.CassandraRead.HealthCheck(ctx)
}

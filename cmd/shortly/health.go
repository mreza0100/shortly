package main

import (
	"context"
	"time"

	"github.com/mreza0100/shortly/cmd/providers"
	"github.com/mreza0100/shortly/internal/services"
	"github.com/urfave/cli"
)

func (a *actions) healthCheck(c *cli.Context) error {
	cassandraRead, _ := providers.CassandraRepositoryProvider(a.cfg.CassandraConnectionConfigs)

	healthService := services.NewHealthService(&services.HealthServiceOptions{
		CassandraRead: cassandraRead,
	})

	timeout := time.Duration(a.cfg.AppConfigs.HealthCheckTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	isHealthy := healthService.CheckHealth(ctx)
	cancel()

	if !isHealthy {
		return cli.NewExitError("Health check failed", 1)
	}
	return nil
}

package main

import (
	"context"
	"time"

	"github.com/mreza0100/shortly/internal/services"
	"github.com/urfave/cli"
)

func (a *actions) healthCheck(c *cli.Context) error {
	cassandraRead, _ := getCassandraRepo(a.cfg.cassandraConnectionConfigs)

	healthService := services.NewHealthService(cassandraRead)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(a.cfg.appConfigs.HealthCheckTimeout*int(time.Second)))

	isHealthy := healthService.CheckHealth(ctx)
	cancel()

	if !isHealthy {
		return cli.NewExitError("Health check failed", 1)
	}
	return nil
}

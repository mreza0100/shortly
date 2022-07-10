package main

import (
	"github.com/mreza0100/shortly/cmd/providers"
	"github.com/mreza0100/shortly/internal/adapters/seed"
	"github.com/mreza0100/shortly/internal/pkg/jwt"
	"github.com/mreza0100/shortly/internal/ports"
	"github.com/mreza0100/shortly/internal/services"
	"github.com/mreza0100/shortly/pkg/convert"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
	"github.com/urfave/cli"
)

func (a *actions) seed(c *cli.Context) error {
	cassandraRead, cassandraWrite := providers.CassandraRepositoryProvider(a.cfg.CassandraConnectionConfigs)
	jwtUtils := jwt.New(a.cfg.AppConfigs.JWTSecret, convert.HourToDuration(a.cfg.AppConfigs.JWTExpire))
	passwordHasher := password_hasher.New(a.cfg.AppConfigs.Salt)
	kgs, err := providers.KGSProvider(cassandraRead, cassandraWrite)
	if err != nil {
		return err
	}

	userService := services.NewUserService(&services.UserServiceDep{
		StorageRead:    cassandraRead,
		StorageWrite:   cassandraWrite,
		JwtUtils:       jwtUtils,
		PasswordHasher: passwordHasher,
	})
	linkService := services.NewLinkService(&services.LinkServiceDep{
		StorageRead:  cassandraRead,
		StorageWrite: cassandraWrite,
		KGS:          kgs,
	})

	seed.SeedDatabase(&ports.Services{
		User: userService,
		Link: linkService,
	})

	return nil
}

package main

import (
	"github.com/mreza0100/shortly/cmd/providers"
	http "github.com/mreza0100/shortly/internal/adapters/http"
	"github.com/mreza0100/shortly/internal/pkg/jwt"
	"github.com/mreza0100/shortly/internal/ports"
	"github.com/mreza0100/shortly/internal/services"
	"github.com/mreza0100/shortly/pkg/convert"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
	"github.com/urfave/cli"
)

func (a *actions) run(c *cli.Context) error {
	cassandraRead, cassandraWrite := providers.GetCassandraRepo(a.cfg.CassandraConnectionConfigs)
	jwtUtils := jwt.New(a.cfg.AppConfigs.JWTSecret, convert.HourToDuration(a.cfg.AppConfigs.JWTExpire))
	passwordHasher := password_hasher.New(a.cfg.AppConfigs.Salt)
	kgs, err := providers.NewKGS(cassandraRead, cassandraWrite)
	if err != nil {
		return err
	}

	userService := services.NewUserService(services.UserServiceOptions{
		CassandraRead:  cassandraRead,
		CassandraWrite: cassandraWrite,
		JwtUtils:       jwtUtils,
		PasswordHasher: passwordHasher,
	})
	linkService := services.NewLinkService(services.LinkServiceOptions{
		CassandraRead:  cassandraRead,
		CassandraWrite: cassandraWrite,
		KGS:            kgs,
	})

	server := http.NewHttpServer(http.NewHttpServerOpts{
		Port:     a.cfg.AppConfigs.Port,
		IsDev:    a.cfg.AppConfigs.IsDev,
		JwtUtils: jwtUtils,
		Services: &ports.Services{
			User: userService,
			Link: linkService,
		},
	})

	return <-server.ListenAndServe()
}

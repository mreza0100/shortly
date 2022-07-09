package main

import (
	http "github.com/mreza0100/shortly/internal/adapters/http"
	"github.com/mreza0100/shortly/internal/pkg/jwt"
	"github.com/mreza0100/shortly/internal/ports"
	"github.com/mreza0100/shortly/internal/services"
	"github.com/mreza0100/shortly/pkg/convert"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
	"github.com/urfave/cli"
)

func (a *actions) run(c *cli.Context) error {
	cassandraRead, cassandraWrite := getCassandraRepo(a.cfg.cassandraConnectionConfigs)
	jwtUtils := jwt.New(a.cfg.appConfigs.JWTSecret, convert.HourToDuration(a.cfg.appConfigs.JWTExpire))
	passwordHasher := password_hasher.New(a.cfg.appConfigs.Salt)
	kgs, err := newKGS(cassandraRead, cassandraWrite)
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
		Port:     a.cfg.appConfigs.Port,
		IsDev:    a.cfg.appConfigs.IsDev,
		JwtUtils: jwtUtils,
		Services: &ports.Services{
			User: userService,
			Link: linkService,
		},
	})

	return <-server.ListenAndServe()
}

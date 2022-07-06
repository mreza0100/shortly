package main

import (
	"log"

	cassandrarepo "github.com/mreza0100/shortly/internal/driven/cassandra"
	"github.com/mreza0100/shortly/internal/ports/driven"
	services_ports "github.com/mreza0100/shortly/internal/ports/services"
	"github.com/mreza0100/shortly/internal/presenters/http"
	"github.com/mreza0100/shortly/internal/services"
	cassandraconnection "github.com/mreza0100/shortly/pkg/connection/cassandra"
	"github.com/mreza0100/shortly/pkg/jwt"
	passwordhasher "github.com/mreza0100/shortly/pkg/password"
	"github.com/urfave/cli"
)

func getCassandraRepo(cfg *Configs) (driven.CassandraReadPort, driven.CassandraWritePort) {
	session, err := cassandraconnection.GetCassandraConnection(&cassandraconnection.ConnectionConfigs{
		Host:     cfg.Cassandra.Host,
		Port:     cfg.Cassandra.Port,
		Keyspace: cfg.Cassandra.Keyspace,
	})
	if err != nil {
		log.Fatal(err)
	}

	cassandraRead, cassandraWrite, err := cassandrarepo.NewCassandraRepository(session)
	if err != nil {
		log.Fatal(err)
	}

	return cassandraRead, cassandraWrite
}

func (a *actions) run(c *cli.Context) error {
	cassandraRead, cassandraWrite := getCassandraRepo(a.cfg)
	jwtUtil := jwt.New(a.cfg.JWTSecret, jwt.HourToDuration(a.cfg.JWTExpire))
	passwordHasher := passwordhasher.New(a.cfg.Salt)

	userService := services.NewUserService(services.UserServiceOptions{
		CassandraRead:  cassandraRead,
		CassandraWrite: cassandraWrite,
		JwtUtil:        jwtUtil,
		PasswordHasher: passwordHasher,
	})

	server := http.NewHttpServer(a.cfg.Port, a.cfg.IsDev, &services_ports.Services{
		User: userService,
	})

	return <-server.ListenAndServe()
}

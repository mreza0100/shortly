package main

import (
	"context"
	"log"

	cassandra_repo "github.com/mreza0100/shortly/internal/adapters/cassandra"
	http "github.com/mreza0100/shortly/internal/adapters/http"
	"github.com/mreza0100/shortly/internal/adapters/kgs"
	"github.com/mreza0100/shortly/internal/ports"
	"github.com/mreza0100/shortly/internal/services"
	cassandra_connection "github.com/mreza0100/shortly/pkg/connection/cassandra"
	"github.com/mreza0100/shortly/pkg/convert"
	"github.com/mreza0100/shortly/pkg/jwt"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
	"github.com/urfave/cli"
)

func getCassandraRepo(cfg *cassandraConnectionConfigs) (ports.CassandraReadPort, ports.CassandraWritePort) {
	session, err := cassandra_connection.CreateConnection(&cassandra_connection.ConnectionConfigs{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Keyspace: cfg.Keyspace,
	})
	if err != nil {
		log.Fatal(err)
	}

	cassandraRead, cassandraWrite, err := cassandra_repo.NewCassandraRepository(session)
	if err != nil {
		log.Fatal(err)
	}

	return cassandraRead, cassandraWrite
}

func newKGS(cassandraRead ports.CassandraReadPort, cassandraWrite ports.CassandraWritePort) (ports.KGS, error) {
	counter, err := cassandraRead.GetCounter(context.Background())
	if err != nil {
		return nil, err
	}

	kgs := kgs.New(kgs.InitKGSOptions{
		SaveCounter: func(newCounter int64) {
			if err := cassandraWrite.UpdateCounter(context.Background(), newCounter); err != nil {
				log.Fatal("Failed to save counter: ", err)
			}
		},
		LastSavedCounter: counter,
	})

	return kgs, nil
}

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

package main

import (
	"context"
	"log"

	cassandra_repo "github.com/mreza0100/shortly/internal/adapters/driven/cassandra"
	"github.com/mreza0100/shortly/internal/adapters/driven/kgs"
	http "github.com/mreza0100/shortly/internal/adapters/driving/http"
	"github.com/mreza0100/shortly/internal/ports/driven"
	services_ports "github.com/mreza0100/shortly/internal/ports/services"
	"github.com/mreza0100/shortly/internal/services"
	cassandra_connection "github.com/mreza0100/shortly/pkg/connection/cassandra"
	"github.com/mreza0100/shortly/pkg/convert"
	"github.com/mreza0100/shortly/pkg/jwt"
	password_hasher "github.com/mreza0100/shortly/pkg/password"
	"github.com/urfave/cli"
)

func getCassandraRepo(cfg *Configs) (driven.CassandraReadPort, driven.CassandraWritePort) {
	session, err := cassandra_connection.CreateConnection(&cassandra_connection.ConnectionConfigs{
		Host:     cfg.Cassandra.Host,
		Port:     cfg.Cassandra.Port,
		Keyspace: cfg.Cassandra.Keyspace,
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

func newKGS(cassandraRead driven.CassandraReadPort, cassandraWrite driven.CassandraWritePort) (kgs.KGS, error) {
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
		LastModifiedCounter: counter,
	})

	return kgs, nil
}

func (a *actions) run(c *cli.Context) error {
	cassandraRead, cassandraWrite := getCassandraRepo(a.cfg)
	jwtUtils := jwt.New(a.cfg.JWTSecret, convert.HourToDuration(a.cfg.JWTExpire))
	passwordHasher := password_hasher.New(a.cfg.Salt)
	kgs, err := newKGS(cassandraRead, cassandraWrite)
	if err != nil {
		return err
	}

	userService := services.NewUserService(services.UserServiceOptions{
		CassandraRead:  cassandraRead,
		CassandraWrite: cassandraWrite,
		JwtUtil:        jwtUtils,
		PasswordHasher: passwordHasher,
	})
	linkService := services.NewLinkService(services.LinkServiceOptions{
		CassandraRead:  cassandraRead,
		CassandraWrite: cassandraWrite,
		KGS:            kgs,
		BaseURL:        a.cfg.Address,
	})

	server := http.NewHttpServer(http.NewHttpServerOpts{
		Port:     a.cfg.Port,
		IsDev:    a.cfg.IsDev,
		JwtUtils: jwtUtils,
		Services: &services_ports.Services{
			User: userService,
			Link: linkService,
		},
	})

	return <-server.ListenAndServe()
}

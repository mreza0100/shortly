package providers

import (
	"context"
	"log"

	cassandra_repo "github.com/mreza0100/shortly/internal/adapters/cassandra"
	cassandra_connection "github.com/mreza0100/shortly/internal/pkg/connections/cassandra"

	"github.com/mreza0100/shortly/internal/adapters/kgs"
	"github.com/mreza0100/shortly/internal/ports"
)

func CassandraRepositoryProvider(cfg *CassandraConnectionConfigs) (ports.StorageReadPort, ports.StorageWritePort) {
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

func KGSProvider(cassandraRead ports.StorageReadPort, cassandraWrite ports.StorageWritePort) (ports.KGS, error) {
	counter, err := cassandraRead.GetCounter(context.Background())
	if err != nil {
		return nil, err
	}

	kgs := kgs.New(&kgs.InitKGSDep{
		SaveCounter: func(newCounter int64) {
			if err := cassandraWrite.UpdateCounter(context.Background(), newCounter); err != nil {
				log.Fatal("Failed to save counter: ", err)
			}
		},
		LastSavedCounter: counter,
	})

	return kgs, nil
}

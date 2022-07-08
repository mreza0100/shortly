package cassandrarepo

import (
	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/adapters/cassandra/migrate"
	"github.com/mreza0100/shortly/internal/ports"
)

func NewCassandraRepository(session *gocql.Session) (ports.CassandraReadPort, ports.CassandraWritePort, error) {
	if err := migrate.MigrateCassandra(session); err != nil {
		return nil, nil, err
	}

	read := &cassandraRead{session: session}
	write := &cassandraWrite{session: session}

	return read, write, nil
}

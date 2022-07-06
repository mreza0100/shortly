package cassandrarepo

import (
	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/driven/cassandra/migrate"
	"github.com/mreza0100/shortly/internal/ports/driven"
)

func NewCassandraRepository(session *gocql.Session) (driven.CassandraReadPort, driven.CassandraWritePort, error) {
	if err := migrate.MigrateCassandra(session); err != nil {
		return nil, nil, err
	}

	read, write := &cassandraRead{session: session}, &cassandraWrite{session: session}

	return read, write, nil
}

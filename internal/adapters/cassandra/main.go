package cassandrarepo

import (
	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/ports"
)

func NewCassandraRepository(session *gocql.Session) (ports.StorageReadPort, ports.StorageWritePort, error) {
	read := &cassandraRead{session: session}
	write := &cassandraWrite{session: session}

	return read, write, nil
}

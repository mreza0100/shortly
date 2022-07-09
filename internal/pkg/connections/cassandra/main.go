package cassandraconnection

import (
	"time"

	"github.com/gocql/gocql"
)

type ConnectionConfigs struct {
	Host     string
	Port     int
	Keyspace string
}

func CreateConnection(cfg *ConnectionConfigs) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.Host)

	cluster.Keyspace = cfg.Keyspace
	cluster.ConnectTimeout = time.Second * 5
	cluster.Timeout = time.Second * 5

	return cluster.CreateSession()
}

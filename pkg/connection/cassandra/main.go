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

func GetCassandraConnection(cfg *ConnectionConfigs) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.Host)
	cluster.Keyspace = cfg.Keyspace
	cluster.ConnectTimeout, cluster.Timeout = time.Second*5, time.Second*5

	return cluster.CreateSession()
}

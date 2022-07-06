package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type CassandraConnectionConfigs struct {
	Host     string
	Port     int
	Keyspace string
}

type Configs struct {
	Address   string
	Port      string
	IsDev     bool
	JWTSecret string
	JWTExpire int
	Salt      string

	Cassandra CassandraConnectionConfigs
}

func getPathFromArgs() string {
	if len(os.Args) > 2 {
		return os.Args[2]
	}
	return ""
}

func loadConfigs() *Configs {
	path := getPathFromArgs()
	if path == "" {
		path = "configs/"
	}

	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("shortly")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	cfg := new(Configs)
	if err := v.Unmarshal(cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

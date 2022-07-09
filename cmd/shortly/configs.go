package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type appConfigs struct {
	Port               string `mapstructure:"app_port"`
	IsDev              bool   `mapstructure:"is_dev"`
	JWTSecret          string `mapstructure:"jwt_secret"`
	JWTExpire          int    `mapstructure:"jwt_expire_hour"`
	Salt               string `mapstructure:"salt"`
	HealthCheckTimeout int    `mapstructure:"health_check_timeout"`
}

type cassandraConnectionConfigs struct {
	Host     string `mapstructure:"cassandra_host"`
	Port     int    `mapstructure:"cassandra_port"`
	Keyspace string `mapstructure:"cassandra_keyspace"`
}

type configs struct {
	appConfigs                 *appConfigs
	cassandraConnectionConfigs *cassandraConnectionConfigs
}

func loadConfigs() *configs {
	v := viper.New()

	var result map[string]interface{}
	var appConfig appConfigs
	var cassandraConfigs cassandraConnectionConfigs
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	if err := v.Unmarshal(&result); err != nil {
		fmt.Printf("Unable to decode into map, %v", err)
	}

	if err := mapstructure.WeakDecode(result, &appConfig); err != nil {
		log.Fatal(err)
	}
	if err := mapstructure.WeakDecode(result, &cassandraConfigs); err != nil {
		log.Fatal(err)
	}

	return &configs{
		appConfigs:                 &appConfig,
		cassandraConnectionConfigs: &cassandraConfigs,
	}
}

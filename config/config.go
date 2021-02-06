package config

import (
	"fmt"
	"os"
	"reflect"
)

type (
	// APIConfig defines the configuration for the api
	APIConfig struct {
		Port string
	}

	// PGConfig defines the configuration for the database
	PGConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
)

// Configuration from environment variables
var (
	API APIConfig
	PG  PGConfig
)

func init() {
	loadConfig()
}

func loadConfig() {
	API = APIConfig{
		Port: os.Getenv("API_PORT"),
	}

	PG = PGConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		DBName:   os.Getenv("PG_DBNAME"),
	}
}

// SanityCheck will panic if there is a missing configuration value
func SanityCheck() {
	cfg := []interface{}{
		API,
		PG,
	}

	for _, v := range cfg {
		rv := reflect.ValueOf(v)
		for i := 0; i < rv.NumField(); i++ {
			if rv.Field(i).String() == "" {
				panic(fmt.Sprintf("[CONFIG]::[%s is missing]", rv.Type().Field(i).Name))
			}
		}
	}
}

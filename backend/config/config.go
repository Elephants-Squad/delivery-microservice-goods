package config

import (
	"flag"
)

type Config struct {
	HTTP
	Log
	DB
}

type HTTP struct {
	Port string
}

type Log struct {
	LogLevel string
}

type DB struct {
	URL string
}

func New() (*Config, error) {
	port := flag.String("port", "8080", "http port")
	logLevel := flag.String("log_level", "local", "log level")
	dbURL := flag.String("db_url", "", "http url")

	flag.Parse()

	cfg := &Config{
		HTTP: HTTP{
			*port,
		},
		Log: Log{
			*logLevel,
		},
		DB: DB{
			*dbURL,
		},
	}

	return cfg, nil
}

package config

import "flag"

type (
	Config struct {
		HTTP
		Log
		DB
	}

	HTTP struct {
		Port string
	}

	DB struct {
		URL string
	}
)

func NewConfig() (*Config, error) {
	port := flag.String("port", "localhost:8080", "port")
	dbURL := flag.String("db_url", "", "url for connection to database")

	flag.Parse()

	cfg := &Config{
		HTTP: HTTP{
			Port: *port,
		},
		DB: DB{
			URL: *dbURL,
		},
	}

	return cfg, nil
}

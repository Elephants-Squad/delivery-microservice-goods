package config

import "flag"

type (
	Config struct {
		HTTP
		DB
		LOG_LEVEL
	}

	HTTP struct {
		Port string
	}

	DB struct {
		URL string
	}

	LOG_LEVEL struct {
		Level string
	}
)

func NewConfig() (*Config, error) {
	port := flag.String("port", "localhost:8080", "port")
	dbURL := flag.String("db_url", "", "url for connection to database")
	logLevel := flag.String("log_level", "info", "log level")

	flag.Parse()

	cfg := &Config{
		HTTP: HTTP{
			Port: *port,
		},
		DB: DB{
			URL: *dbURL,
		},
		LOG_LEVEL: LOG_LEVEL{
			Level: *logLevel,
		},
	}

	return cfg, nil
}

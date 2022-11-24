package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Logger struct {
	Level    string `env:"LOG_LEVEL" envDefault:"warn"`
	Encoding string `env:"LOG_ENCODING" envDefault:"json"`
}

type Config struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"golangUnitedSchool"`
	Port        string `env:"PORT" envDefault:"5000"`
	Host        string `env:"HOST" envDefault:"0.0.0.0"`
	PgDsn       string `env:"PG_DSN" envDefault:"postgres://pguser:pguser@localhost:5432/pgdb?sslmode=disable"`
	Logger      Logger
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("can not load config: %s", err)
	}
	return &cfg, nil
}

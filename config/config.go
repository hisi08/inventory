package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "inventory"
)

type Config struct {
	Database     Database
	Server       Server
	DBConnection DBConnection
}

type Database struct {
	PostgresURL string `env:"POSTGRESURL" envDefault:""`
	Name        string `env:"DB_NAME" envDefault:"inventory"`
	Port        int    `env:"DB_PORT" envDefault:"5432"`
	User        string `env:"DB_user" envDefault:"postgres"`
}

type DBConnection struct {
	Host     string `env:"host" envDefault:"localhost"`
	Port     int    `env:"port" envDefault:"5432"`
	User     string `env:"user" envDefault:"postgres"`
	Password string `env:"password" envDefault:"root"`
	Dbname   string `env:"dbname" envDefault:"inventory"`
}

type Server struct {
	GRPCHost string `env:"GRPCHost" envDefault:"0.0.0.0"`
	GRPCPort int    `env:"GRPCPort" envDefault:"9090"`
}

func Load() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return &cfg, errors.Wrap(err, "failed to load environment")
	}
	return &cfg, nil
}

func GetPostgresURL() string {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	psqlURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		user, password, host, port, dbname)

	return psqlURL
}

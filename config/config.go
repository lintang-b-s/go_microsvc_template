package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App
		HTTP
		Postgres
		LogConfig
		GRPC
		// Auth

	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Port string `env-required:"true" env:"HTTP_PORT"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}


	Postgres struct {
		Username string `env-required:"true" yaml:"username"  env:"USERNAME_POSTGRES"`
		PGURL    string `json:"pg_url" yaml:"pg_url" env:"PG_URL"`
		Password string `env-required:"true" yaml:"password" env:"PASSWORD_POSTGRES"`
		PGScheme string `env-required:"true" json:"pg_scheme" yaml:"pg_scheme" env:"PG_SCHEME"`
		PGDB     string `env-required:"true" json:"pg_db" yaml:"pg_db" env:"PG_DB"`
	}

	LogConfig struct {
		Level string `json:"level" yaml:"level" env:"LOG_LEVEL"`
		// Filename   string `json:"filename" yaml:"filename"`
		// MaxSize    int    `json:"maxsize" yaml:"maxsize"`
		MaxAge     int `json:"max_age" yaml:"max_age" env:"LOG_MAXAGE"`
		MaxBackups int `json:"max_backups" yaml:"max_backups" env:"LOG_MAXBACKUP"`
	}

	GRPC struct {
		URLGrpc    string `json:"urlGRPC" yaml:"urlGRPC" env:"URL_GRPC"`
	}

	// Auth struct {
	// 	PrivateKey string `json:"private_key" env:"PRIVATE_KEY"`
	// 	PublicKey string `json:"public_key" env:"PUBLIC_KEY"`
	// }
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// err = cleanenv.ReadConfig(path+".env", cfg) // buat di doker , ../.env kalo debug (.env kalo docker)
	// err = cleanenv.ReadConfig(path+"/local.env", cfg) // local run
	if os.Getenv("APP_ENV") == "local" {
		err = cleanenv.ReadConfig(path+"/local.env", cfg)
	} else {
		err = cleanenv.ReadConfig(path+".env", cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

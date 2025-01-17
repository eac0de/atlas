package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	// Common
	ProjectName string `env:"PROJECT_NAME" envDefault:"auth"`
	IsDev       bool   `env:"IS_DEV" envDefault:"true"`

	// Server
	ServerAddress     string `env:"SERVER_ADDRESS" envDefault:"0.0.0.0:8080"`
	GPRCServerAddress string `env:"GRPC_SERVER_ADDRESS" envDefault:"0.0.0.0:9090"`

	// PSQL
	PSQLHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PSQLPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	PSQLUsername string `env:"POSTGRES_USER" envDefault:"postgres"`
	PSQLPassword string `env:"POSTGRES_PASSWORD" envDefault:"351762"`
	PSQLDBName   string `env:"POSTGRES_DB" envDefault:"sentinel"`

	// JWT
	JWTSecretKey  string        `env:"JWT_SECRET_KEY" envDefault:"supersecretkey"`
	JWTAccessExp  time.Duration `env:"JWT_ACCESS_EXP" envDefault:"15m"`
	JWTRefreshExp time.Duration `env:"JWT_REFRESH_EXP" envDefault:"168h"`

	// EMAIL_SENDER
	SMTPHost     string `env:"SMTP_HOST"`
	SMTPPort     string `env:"SMTP_PORT"`
	SMTPUsername string `env:"SMTP_USERNAME"`
	SMTPPassword string `env:"SMTP_PASSWORD"`
}

func MustLoad() *Config {
	cfg := new(Config)
	err := env.Parse(cfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return cfg
}

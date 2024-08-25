package config

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DBUser   string `env:"DB_USER, required"`
	DBPasswd string `env:"DB_PASSWD, required"`
	DBAddr   string `env:"DB_ADDR, required"`
	DBName   string `env:"DB_NAME, required"`
}

func New(ctx context.Context) (*Config, error) {
	var conf Config
	if err := envconfig.Process(ctx, &conf); err != nil {
		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	return &Config{
		DBUser:   conf.DBUser,
		DBPasswd: conf.DBPasswd,
		DBAddr:   conf.DBAddr,
		DBName:   conf.DBName,
	}, nil
}

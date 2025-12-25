package config

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	SpannerProject  string `env:"SPANNER_PROJECT, required"`
	SpannerInstance string `env:"SPANNER_INSTANCE, required"`
	SpannerDatabase string `env:"SPANNER_DATABASE, required"`
	SpannerEmulator string `env:"SPANNER_EMULATOR_HOST, default="`
}

func New(ctx context.Context) (*Config, error) {
	var conf Config
	if err := envconfig.Process(ctx, &conf); err != nil {
		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	return &conf, nil
}

func (c *Config) DatabaseName() string {
	return "projects/" + c.SpannerProject + "/instances/" + c.SpannerInstance + "/databases/" + c.SpannerDatabase
}

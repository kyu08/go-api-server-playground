package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/config"
	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

type MarketServer struct {
	api.UnimplementedMarketServiceServer
}

func NewMarketServer() (*MarketServer, error) {
	config, err := config.New(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = database.NewDBConnection(config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &MarketServer{
		UnimplementedMarketServiceServer: api.UnimplementedMarketServiceServer{},
	}, nil
}

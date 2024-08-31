package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kyu08/go-api-server-playground/internal/config"
	"github.com/kyu08/go-api-server-playground/internal/handler"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	sloggin "github.com/samber/slog-gin"
)

func main() {
	const (
		host = ""
		port = 8080
	)

	config, err := config.New(context.Background())
	if err != nil {
		panic(err)
	}

	db, err := database.NewDBConnection(config)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router.Use(gin.Recovery(), sloggin.New(logger))

	router.GET("/health", handler.Health)
	router.POST("/sample", handler.Sample(usecase.NewSampleUsecase(db, repository.NewSampleRepository())))

	if err := router.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
		// TODO: graceful shutdown
		panic(err)
	}
}

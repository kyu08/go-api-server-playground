package usecase

import (
	"context"
	"database/sql"

	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/repository"
)

type (
	SampleUsecase struct {
		db               *sql.DB
		sampleRepository *repository.SampleRepository
	}
	SampleInput struct {
		Name string
	}
	SampleOutput struct {
		Message string
	}
)

func (u SampleUsecase) Run(ctx context.Context, input *SampleInput) (*SampleOutput, error) {
	return &SampleOutput{
		Message: "Hello, " + input.Name,
	}, nil
}

func NewSampleUsecase(db *sql.DB, sampleRepository *repository.SampleRepository) *SampleUsecase {
	return &SampleUsecase{
		db:               db,
		sampleRepository: sampleRepository,
	}
}

func NewSampleInput(name string) *SampleInput {
	return &SampleInput{
		Name: name,
	}
}

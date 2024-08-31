package repository

import (
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
)

type SampleRepository struct {
	queries *database.Queries
}

func NewSampleRepository() *SampleRepository {
	return &SampleRepository{
		queries: nil,
	}
}

func (r *SampleRepository) SetQueries(q *database.Queries) {
	r.queries = q
}

package service

import (
	"context"

	sensor "github.com/ghanto/sds011-server/entity"
	"github.com/pkg/errors"
)

type SensorRepository interface {
	Add(ctx context.Context, record sensor.Record) error
	Get(ctx context.Context) ([]sensor.Record, error)
}

type SdsService struct {
	repo SensorRepository
}

// NewSdsService creates a new instance of SDS011 service
func NewSdsService(repo SensorRepository) *SdsService {
	return &SdsService{
		repo: repo,
	}
}

func (s *SdsService) Add(ctx context.Context, record sensor.Record) error {
	if err := s.repo.Add(ctx, record); err != nil {
		errorMessage := "Unable to save record"
		return errors.Wrap(err, errorMessage)
	}
	return nil
}

func (s *SdsService) Get(ctx context.Context) ([]sensor.Record, error) {
	records, err := s.repo.Get(ctx)

	if err != nil {
		errorMessage := "Unable to get records"
		return []sensor.Record{}, errors.Wrap(err, errorMessage)
	}

	return records, nil
}

package service

import (
	"context"

	sensor "github.com/ghanto/sds011-server/entity"
	"github.com/pkg/errors"
)

// SensorRepository interface signature
type SensorRepository interface {
	Add(ctx context.Context, record sensor.Record) error
	Get(ctx context.Context) ([]sensor.Record, error)
}

// SdsService service signature
type SdsService struct {
	repo SensorRepository
}

// NewSdsService creates a new instance of SDS011 service
func NewSdsService(repo SensorRepository) *SdsService {
	return &SdsService{
		repo: repo,
	}
}

// Add an item to the repository
func (s *SdsService) Add(ctx context.Context, record sensor.Record) error {
	if err := s.repo.Add(ctx, record); err != nil {
		errorMessage := "Unable to save record"
		return errors.Wrap(err, errorMessage)
	}
	return nil
}

// Get get all items from repository
func (s *SdsService) Get(ctx context.Context) ([]sensor.Record, error) {
	records, err := s.repo.Get(ctx)

	if err != nil {
		errorMessage := "Unable to get records"
		return []sensor.Record{}, errors.Wrap(err, errorMessage)
	}

	return records, nil
}

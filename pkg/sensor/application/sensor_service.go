package application

import (
	"context"
	"fmt"

	sensor "github.com/ghanto/sds011-server/pkg/sensor/domain"
	"github.com/pkg/errors"
)

// SensorRepository interface signature
type SensorRepository interface {
	Add(ctx context.Context, record []byte) error
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
	serialized := s.serialize(ctx, record)

	if err := s.repo.Add(ctx, serialized); err != nil {
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

func (s *SdsService) serialize(ctx context.Context, record sensor.Record) []byte {
	serialized := []byte(fmt.Sprintf("%v", record))
	return serialized
}

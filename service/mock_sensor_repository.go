package service

import (
	"context"

	sensor "github.com/ghanto/sds011-server/entity"
)

// MockSensorRepository mock repository structure
type MockSensorRepository struct {
	records []sensor.Record
	err     error
}

// NewMockRepository creates a new instance of mock repository
func NewMockRepository(r []sensor.Record, e error) *MockSensorRepository {
	return &MockSensorRepository{
		records: r,
		err:     e,
	}
}

// Add an item to the repository
func (s *MockSensorRepository) Add(ctx context.Context, record []byte) error {
	return s.err
}

// Get get all items from repository
func (s *MockSensorRepository) Get(ctx context.Context) ([]sensor.Record, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.records, nil
}

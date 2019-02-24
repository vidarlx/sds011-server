package service

import (
	"context"

	sensor "github.com/ghanto/sds011-server/entity"
)

type MockSensorRepository struct {
	records []sensor.Record
	err     error
}

func NewMockRepository(r []sensor.Record, e error) *MockSensorRepository {
	return &MockSensorRepository{
		records: r,
		err:     e,
	}
}

func (s *MockSensorRepository) Add(ctx context.Context, record sensor.Record) error {
	return s.err
}

func (s *MockSensorRepository) Get(ctx context.Context) ([]sensor.Record, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.records, nil
}

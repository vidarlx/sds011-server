package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	sensor "github.com/ghanto/sds011-server/entity"
	"github.com/pkg/errors"
)

func TestSdsService_Add(t *testing.T) {
	testCases := map[string]struct {
		record        sensor.Record
		createError   error
		expectedError error
	}{
		"evertyhing fine": {
			record:        sensor.Record{},
			createError:   nil,
			expectedError: nil,
		},
		"error, repo should return error": {
			record:        sensor.Record{},
			createError:   errors.New("panic error"),
			expectedError: errors.New("panic error"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			repo := NewMockRepository([]sensor.Record{}, tc.createError)
			sensorService := NewSdsService(repo)
			if err := sensorService.Add(context.Background(), tc.record); err != nil {
				errCause := errors.Cause(err)
				if fmt.Sprintf("%v", errCause) != fmt.Sprintf("%v", tc.expectedError) {
					t.Errorf("expected error=%v got %v", tc.expectedError, errCause)
				}
			}

		})
	}
}

func TestSdsService_Get(t *testing.T) {
	mockRecords := []sensor.Record{
		sensor.Record{
			PM10: 12,
			PM25: 14,
		},
		sensor.Record{
			PM10: 15,
			PM25: 20,
		},
	}

	testCases := map[string]struct {
		expectedRecordsLength int
		createError           error
		expectedError         error
	}{
		"evertyhing fine": {
			expectedRecordsLength: 2,
			createError:           nil,
			expectedError:         nil,
		},
		"error, repo should return error": {
			expectedRecordsLength: 0,
			createError:           errors.New("panic error"),
			expectedError:         errors.New("panic error"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			repo := NewMockRepository(mockRecords, tc.createError)
			sensorService := NewSdsService(repo)

			records, err := sensorService.Get(context.Background())
			fmt.Println(records)
			if err != nil {
				errCause := errors.Cause(err)
				if fmt.Sprintf("%v", errCause) != fmt.Sprintf("%v", tc.expectedError) {
					t.Errorf("expected error=%v got %v", tc.expectedError, errCause)
				}
			}

			assert.Len(t, records, tc.expectedRecordsLength)
		})
	}
}

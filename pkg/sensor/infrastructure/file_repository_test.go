package infrastructure

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
)

func TestFileRepository_Write(t *testing.T) {
	testCases := map[string]struct {
		bodyFile      []byte
		writeError    error
		readError     error
		expectedError error
		expectedBody  []byte
	}{
		"everything fine": {
			bodyFile:      []byte(`test`),
			expectedBody:  []byte(`test`),
			writeError:    nil,
			readError:     nil,
			expectedError: nil,
		},
		"save error, repo should return error": {
			bodyFile:      nil,
			writeError:    errors.New("fs is gone"),
			expectedError: errors.New("fs is gone"),
		},
		"read error, repo should return error": {
			bodyFile:      []byte(`test`),
			readError:     errors.New("hard drive is gone"),
			expectedError: errors.New("hard drive is gone"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			frMock := func(filename string) ([]byte, error) {
				return tc.expectedBody, tc.readError
			}

			fwMock := func(filename string, data []byte, perm os.FileMode) error {
				return tc.writeError
			}
			fileRepo := NewFileRepository(frMock, fwMock)

			err := fileRepo.Save(`test.txt`, tc.bodyFile)
			if err != nil {
				errCause := errors.Cause(err)
				if fmt.Sprintf("%v", errCause) != fmt.Sprintf("%v", tc.expectedError) {
					t.Errorf("expected error=%v got %v", tc.expectedError, errCause)
				}
			}

			body, err := fileRepo.Load(`test.txt`)
			if err != nil {
				errCause := errors.Cause(err)
				if fmt.Sprintf("%v", errCause) != fmt.Sprintf("%v", tc.expectedError) {
					t.Errorf("expected error=%v got %v", tc.expectedError, errCause)
				}
			}

			assert.Equal(t, tc.expectedBody, body)
		})
	}
}

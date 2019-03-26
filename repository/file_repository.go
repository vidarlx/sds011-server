package repository

import (
	"os"

	"github.com/pkg/errors"
)

type fileReader = func(filename string) ([]byte, error)
type fileWriter = func(filename string, data []byte, perm os.FileMode) error

// FileRepository represents file repository
type FileRepository struct {
	Reader fileReader
	Writer fileWriter
}

// NewFileRepository create a new file repository
func NewFileRepository(r fileReader, w fileWriter) *FileRepository {
	return &FileRepository{
		Reader: r,
		Writer: w,
	}
}

// Save saves a record to a file
func (r *FileRepository) Save(filename string, record []byte) error {
	if err := r.Writer(filename, record, 0644); err != nil {
		return errors.Wrap(err, "file repository can't save file")
	}
	return nil
}

// Load reads a records from a file
func (r *FileRepository) Load(filename string) ([]byte, error) {
	record, err := r.Reader(filename)
	if err != nil {
		return nil, errors.Wrap(err, "file repository can't save file")
	}
	return record, nil
}

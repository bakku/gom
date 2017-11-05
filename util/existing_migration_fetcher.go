package util

import (
	"io/ioutil"
	"os"
)

type dirReaderInterface interface {
	Read() ([]os.FileInfo, error)
}

type dirReader struct{}

func (d *dirReader) Read() ([]os.FileInfo, error) {
	dirs, err := ioutil.ReadDir("db/migrations")
	if err != nil {
		return []os.FileInfo{}, err
	}
	return dirs, nil
}

type ExistingMigrationFetcher struct {
	DirReader dirReaderInterface
}

func NewExistingMigrationFetcher() *ExistingMigrationFetcher {
	return &ExistingMigrationFetcher{
		DirReader: &dirReader{},
	}
}

func (e *ExistingMigrationFetcher) Fetch() ([]string, error) {
	dirs, err := e.DirReader.Read()
	if err != nil {
		return []string{}, err
	}

	names := make([]string, len(dirs))

	for i, v := range dirs {
		names[i] = v.Name()
	}

	return names, err
}

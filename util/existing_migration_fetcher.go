package util

import (
	"errors"
)

type ExistingMigrationFetcher struct {
	DirReader  dirReaderInterface
	DirChecker dirCheckerInterface
}

func NewExistingMigrationFetcher() *ExistingMigrationFetcher {
	return &ExistingMigrationFetcher{
		DirReader: &dirReader{},
		DirChecker: &dirChecker{},
	}
}

func (e *ExistingMigrationFetcher) Fetch() ([]string, error) {
	if e.DirChecker.DirExists() == false {
		return []string{}, errors.New("migration directory does not exist")
	}

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

package util

import (
	"errors"
)

type ExistingMigrationFetcher struct {
	DirReader  DirReaderInterface
	DirChecker DirCheckerInterface
}

func NewExistingMigrationFetcher() *ExistingMigrationFetcher {
	return &ExistingMigrationFetcher{
		DirReader:  &DirReader{},
		DirChecker: &DirChecker{},
	}
}

func (e *ExistingMigrationFetcher) Fetch() ([]string, error) {
	if e.DirChecker.DirExists("db/migrations") == false {
		return []string{}, errors.New("migration directory does not exist")
	}

	dirs, err := e.DirReader.Read("db/migrations")
	if err != nil {
		return []string{}, err
	}

	names := make([]string, len(dirs))

	for i, v := range dirs {
		names[i] = v.Name()
	}

	return names, err
}

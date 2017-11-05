package util

import (
	"errors"
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

type dirCheckerInterface interface {
	DirExists() bool
}

type dirChecker struct{}


func (d *dirChecker) DirExists() bool {
	if _, err := os.Stat("db/migrations"); os.IsNotExist(err) {
		return false
	}
	return true
}

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

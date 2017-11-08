package util

import (
	"io/ioutil"
	"os"
)

type DirReaderInterface interface {
	Read() ([]os.FileInfo, error)
}

type DirReader struct{}

func (d *DirReader) Read() ([]os.FileInfo, error) {
	dirs, err := ioutil.ReadDir("db/migrations")
	if err != nil {
		return []os.FileInfo{}, err
	}
	return dirs, nil
}

type DirCheckerInterface interface {
	DirExists() bool
}

type DirChecker struct{}

func (d *DirChecker) DirExists() bool {
	if _, err := os.Stat("db/migrations"); os.IsNotExist(err) {
		return false
	}
	return true
}

type DirCreatorInterface interface {
	DirCreate(path string) error
}

type DirCreator struct{}

func (d *DirCreator) DirCreate(path string) error {
	return os.MkdirAll(path, 0755)
}

type FileCreatorInterface interface {
	FileCreate(path string) error
}

type FileCreator struct{}

func (f *FileCreator) FileCreate(path string) error {
	_, err := os.Create(path)
	return err
}

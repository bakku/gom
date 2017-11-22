package util

import (
	"io/ioutil"
	"os"
)

type DirReaderInterface interface {
	Read(path string) ([]os.FileInfo, error)
}

type DirReader struct{}

func (d *DirReader) Read(path string) ([]os.FileInfo, error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return []os.FileInfo{}, err
	}
	return dirs, nil
}

type DirCheckerInterface interface {
	DirExists(path string) bool
}

type DirChecker struct{}

func (d *DirChecker) DirExists(path string) bool {
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

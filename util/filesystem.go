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


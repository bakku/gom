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

type FileDirCheckerInterface interface {
	FileDirExists(path string) bool
}

type FileDirChecker struct{}

func (d *FileDirChecker) FileDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
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

type FileWriterInterface interface {
	Write(path, content string) error
}

type FileWriter struct{}

func (f *FileWriter) Write(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

type FileAppenderInterface interface {
	Append(path, content string) error
}

type FileAppender struct{}

func (f *FileAppender) Append(path, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err := f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

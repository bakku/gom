package mocks

import (
	"os"
	"time"
)

type FileInfo struct {
	N string
}

func (m *FileInfo) Name() string {
	return m.N
}

func (m *FileInfo) Size() int64 {
	return 0
}

func (m *FileInfo) Mode() os.FileMode {
	return os.FileMode(1)
}

func (m *FileInfo) ModTime() time.Time {
	return time.Now()
}

func (m *FileInfo) IsDir() bool {
	return true
}

func (m *FileInfo) Sys() interface{} {
	return nil
}

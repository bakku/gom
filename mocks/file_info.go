package mocks

import (
	"os"
	"time"
)

type MockFileInfo struct {
	N string
}

func (m *MockFileInfo) Name() string {
	return m.N
}

func (m *MockFileInfo) Size() int64 {
	return 0
}

func (m *MockFileInfo) Mode() os.FileMode {
	return os.FileMode(1)
}

func (m *MockFileInfo) ModTime() time.Time {
	return time.Now()
}

func (m *MockFileInfo) IsDir() bool {
	return true
}

func (m *MockFileInfo) Sys() interface{} {
	return nil
}

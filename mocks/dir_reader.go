package mocks

import (
	"os"
)

type MockDirReader struct {
	ReadCall struct {
		Returns struct {
			DirSlice []os.FileInfo
			Error       error
		}
	}
}

func (m *MockDirReader) Read() ([]os.FileInfo, error) {
	return m.ReadCall.Returns.DirSlice, m.ReadCall.Returns.Error
}

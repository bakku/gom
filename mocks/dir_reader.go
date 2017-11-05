package mocks

import (
	"os"
)

type DirReader struct {
	ReadCall struct {
		Returns struct {
			DirSlice []os.FileInfo
			Error    error
		}
	}
}

func (m *DirReader) Read() ([]os.FileInfo, error) {
	return m.ReadCall.Returns.DirSlice, m.ReadCall.Returns.Error
}

package mocks

type FileReader struct {
	ReadCall struct {
		Receives struct {
			Path []string
		}
		Returns struct {
			String []string
			Error  []error
		}
	}
	CallCount uint
}

func (f *FileReader) Read(path string) (string, error) {
	f.ReadCall.Receives.Path = append(f.ReadCall.Receives.Path, path)

	retStr := f.ReadCall.Returns.String[f.CallCount]
	retErr := f.ReadCall.Returns.Error[f.CallCount]

	f.CallCount++

	return retStr, retErr
}

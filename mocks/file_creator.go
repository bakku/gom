package mocks

type FileCreator struct {
	FileCreateCall struct {
		Returns struct {
			Errors struct {
				OnCall int
				Error  error
			}
		}
	}
	ReceivedPaths []string
	callCounter   int
}

func (f *FileCreator) FileCreate(path string) error {
	f.ReceivedPaths = append(f.ReceivedPaths, path)

	f.callCounter++

	if f.callCounter == f.FileCreateCall.Returns.Errors.OnCall {
		return f.FileCreateCall.Returns.Errors.Error
	}

	return nil
}

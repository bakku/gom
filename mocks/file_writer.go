package mocks

type FileWriter struct {
	WriteCall struct {
		Receives struct {
			Path    string
			Content string
		}
		Returns struct {
			Error error
		}
	}
}

func (f *FileWriter) Write(path, content string) error {
	f.WriteCall.Receives.Path = path
	f.WriteCall.Receives.Content = content

	return f.WriteCall.Returns.Error
}

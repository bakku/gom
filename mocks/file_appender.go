package mocks

type FileAppender struct {
	AppendCall struct {
		Receives struct {
			Path    string
			Content string
		}
		Returns struct {
			Error error
		}
	}
}

func (f *FileAppender) Append(path, content string) error {
	f.AppendCall.Receives.Path = path
	f.AppendCall.Receives.Content = content

	return f.AppendCall.Returns.Error
}

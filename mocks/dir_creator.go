package mocks

type DirCreator struct {
	DirCreateCall struct {
		Receives struct {
			Path string
		}
		Returns struct {
			Error error
		}
	}
}

func (d *DirCreator) DirCreate(path string) error {
	d.DirCreateCall.Receives.Path = path

	return d.DirCreateCall.Returns.Error
}

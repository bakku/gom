package mocks

type FileDirChecker struct {
	FileDirExistsCall struct {
		Receives struct {
			Path string
		}
		Returns struct {
			Bool bool
		}
	}
}

func (d *FileDirChecker) FileDirExists(path string) bool {
	d.FileDirExistsCall.Receives.Path = path
	return d.FileDirExistsCall.Returns.Bool
}

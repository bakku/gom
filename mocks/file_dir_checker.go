package mocks

type FileDirChecker struct {
	FileDirExistsCall struct {
		Returns struct {
			Bool bool
		}
	}
}

func (d *FileDirChecker) FileDirExists(path string) bool {
	return d.FileDirExistsCall.Returns.Bool
}

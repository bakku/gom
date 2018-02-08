package mocks

type FileDirChecker struct {
	FileDirExistsCall struct {
		Receives struct {
			Path []string
		}
		Returns struct {
			Bool []bool
		}
	}
	CallCount uint
}

func (d *FileDirChecker) FileDirExists(path string) bool {
	d.FileDirExistsCall.Receives.Path = append(d.FileDirExistsCall.Receives.Path, path)

	exists := d.FileDirExistsCall.Returns.Bool[d.CallCount]
	d.CallCount++

	return exists
}

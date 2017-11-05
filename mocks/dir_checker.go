package mocks

type DirChecker struct {
	DirExistsCall struct {
		Returns struct {
			Bool bool
		}
	}
}

func (d *DirChecker) DirExists() bool {
	return d.DirExistsCall.Returns.Bool
}

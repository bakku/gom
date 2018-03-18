package commands

import (
	"errors"
	"fmt"

	"github.com/bakku/gom/util"
)

func fetchAvailableMigrations(fileChecker util.FileDirCheckerInterface, dirReader util.DirReaderInterface) ([]string, error) {
	if fileChecker.FileDirExists("db/migrations") == false {
		return nil, errors.New("migrate: migrations directory does not exist")
	}

	dirs, err := dirReader.Read("db/migrations")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("migrate: could not read migrations directory: %v", err))
	}

	names := make([]string, len(dirs))

	for i, v := range dirs {
		names[i] = v.Name()
	}

	return names, nil
}

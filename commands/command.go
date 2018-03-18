package commands

import (
	"errors"
	"fmt"
)

type Command interface {
	Run(args ...string) error
}

func Select(command string) (Command, error) {
	switch command {
	case "generate":
		return NewGenerator(), nil
	case "migrate":
		return NewMigrator()
	case "init":
		return NewInitializer()
	case "rollback":
		return NewBackroller()
	default:
		return nil, errors.New(fmt.Sprintf("the command '%s' does not exist", command))
	}
}

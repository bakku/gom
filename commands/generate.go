package commands

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bakku/gom/util"
)

type Generator struct {
	DirCreator  util.DirCreatorInterface
	FileCreator util.FileCreatorInterface
}

func NewGenerator() *Generator {
	return &Generator{
		&util.DirCreator{},
		&util.FileCreator{},
	}
}

func (g *Generator) Run(args ...string) error {
	if len(args) == 0 {
		return errors.New("generate: no migration name passed")
	}

	currTime := time.Now()
	timeFormatted := currTime.Format("20060102150405")
	migrationDir := timeFormatted + "_" + args[0]

	path := []string{"db", "migrations", migrationDir}

	if err := g.DirCreator.DirCreate(strings.Join(path, "/")); err != nil {
		return errors.New(fmt.Sprintf("generate: could not create directory: %v", err))
	}

	upPath := append(path, "up.sql")
	downPath := append(path, "down.sql")

	if err := g.FileCreator.FileCreate(strings.Join(upPath, "/")); err != nil {
		return errors.New(fmt.Sprintf("generate: could not create file: %v", err))
	}

	if err := g.FileCreator.FileCreate(strings.Join(downPath, "/")); err != nil {
		return errors.New(fmt.Sprintf("generate: could not create file: %v", err))
	}

	return nil
}

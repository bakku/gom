package commands

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bakku/gom/util"
)

const insertStmt = "INSERT INTO schema_migrations VALUES (\"%s\") ;\n"

type Generator struct {
	FileDirChecker util.FileDirCheckerInterface
	FileAppender   util.FileAppenderInterface
	DirCreator     util.DirCreatorInterface
	FileCreator    util.FileCreatorInterface
}

func NewGenerator() *Generator {
	return &Generator{
		&util.FileDirChecker{},
		&util.FileAppender{},
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

	if err := g.appendMigrationToSchema(timeFormatted); err != nil {
		return err
	}

	if err := g.createMigrationFiles(timeFormatted, args[0]); err != nil {
		return err
	}

	return nil
}

func (g *Generator) appendMigrationToSchema(formattedTime string) error {
	path := []string{"db", "schema.sql"}

	if ok := g.FileDirChecker.FileDirExists(strings.Join(path, "/")); !ok {
		return errors.New("generate: schema.sql file does not exist")
	}

	if err := g.FileAppender.Append(strings.Join(path, "/"), fmt.Sprintf(insertStmt, formattedTime)); err != nil {
		return errors.New(fmt.Sprintf("generate: could not append new migration to schema.sql: %v", err))
	}

	return nil
}

func (g *Generator) createMigrationFiles(formattedTime, migrationName string) error {
	migrationDir := formattedTime + "_" + migrationName

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

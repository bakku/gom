package commands

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/bakku/gom/util"
)

var schema string = `
CREATE TABLE IF NOT EXISTS schema_migrations (
	migration CHAR(14)		
);
`

var baseSchemaFile string = `CREATE TABLE schema_migrations (
	migration CHAR(14)
) ;

`

type Initializer struct {
	DB             *sql.DB
	DirCreator     util.DirCreatorInterface
	FileDirChecker util.FileDirCheckerInterface
	FileWriter     util.FileWriterInterface
}

func NewInitializer() (*Initializer, error) {
	db, err := util.InitDB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("init: %v", err))
	}

	return &Initializer{
		db,
		&util.DirCreator{},
		&util.FileDirChecker{},
		&util.FileWriter{},
	}, nil
}

func (i *Initializer) Run(args ...string) error {
	_, err := i.DB.Exec(schema)
	if err != nil {
		return errors.New(fmt.Sprintf("init: could not create schema_migrations table: %v", err))
	}

	if err := i.DirCreator.DirCreate("db"); err != nil {
		return errors.New(fmt.Sprintf("init: could not create db folder: %v", err))
	}

	if err := i.DirCreator.DirCreate("db/migrations"); err != nil {
		return errors.New(fmt.Sprintf("init: could not create migrations folder: %v", err))
	}

	if i.FileDirChecker.FileDirExists("db/schema.sql") == false {
		if err := i.FileWriter.Write("db/schema.sql", baseSchemaFile); err != nil {
			return errors.New(fmt.Sprintf("init: could not create base schema.sql: %v", err))
		}
	}

	return nil
}

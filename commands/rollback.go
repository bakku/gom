package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/bakku/gom/util"
)

const rollbackDeleteStmt = "DELETE FROM schema_migrations WHERE migration = '%s' ;"

type Backroller struct {
	DB             *sql.DB
	FileDirChecker util.FileDirCheckerInterface
	DirReader      util.DirReaderInterface
	FileReader     util.FileReaderInterface
}

func NewBackroller() (*Backroller, error) {
	db, err := util.InitDB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("rollback: %v", err))
	}

	return &Backroller{
		db,
		&util.FileDirChecker{},
		&util.DirReader{},
		&util.FileReader{},
	}, nil
}

func (b *Backroller) Run(args ...string) error {
	migratedMigrations, err := util.GetSchemaMigrations(b.DB)
	if err != nil {
		return errors.New(fmt.Sprintf("rollback: could not get schema migrations from database: %v", err))
	}

	if len(migratedMigrations) == 0 {
		return nil
	}

	availableMigrations, err := fetchAvailableMigrations(b.FileDirChecker, b.DirReader)
	if err != nil {
		return errors.New(fmt.Sprintf("rollback: could not get available migrations: %v", err))
	}

	migrationToRollback := migratedMigrations[len(migratedMigrations)-1]

	var migrationToRollbackPath string

	for _, migration := range availableMigrations {
		if splitted := strings.Split(migration, "_"); splitted[0] == migrationToRollback {
			migrationToRollbackPath = migration
			break
		}
	}

	if migrationToRollbackPath == "" {
		return errors.New(fmt.Sprintf("rollback: migration file does not exist for %s", migrationToRollback))
	}

	fullPath := []string{"db", "migrations", migrationToRollbackPath, "down.sql"}

	rollbackStmt, err := b.FileReader.Read(strings.Join(fullPath, "/"))
	if err != nil {
		return errors.New(fmt.Sprintf("rollback: could not read the migration file: %v", err))
	}

	_, err = b.DB.Exec(rollbackStmt)
	if err != nil {
		return errors.New(fmt.Sprintf("rollback: could not execute the rollback statement: %v", err))
	}

	_, err = b.DB.Exec(fmt.Sprintf(rollbackDeleteStmt, migrationToRollback))
	if err != nil {
		return errors.New(fmt.Sprintf("rollback: could not remove the migration from the schema_migrations table: %v", err))
	}

	return nil
}

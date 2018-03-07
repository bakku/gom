package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/bakku/gom/util"
)

const migrationInsertStmt = "INSERT INTO schema_migrations VALUES ('%s') ;"

type Migrator struct {
	FileDirChecker util.FileDirCheckerInterface
	DirReader      util.DirReaderInterface
	DB             *sql.DB
	FileReader     util.FileReaderInterface
}

func NewMigrator() (*Migrator, error) {
	db, err := util.InitDB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("migrate: %v", err))
	}

	return &Migrator{
		&util.FileDirChecker{},
		&util.DirReader{},
		db,
		&util.FileReader{},
	}, nil
}

func (m *Migrator) Run(args ...string) error {
	availableMigrations, err := m.fetchAvailableMigrations()
	if err != nil {
		return err
	}

	migratedMigrations, err := util.GetSchemaMigrations(m.DB)
	if err != nil {
		return err
	}

	migrations := m.getMigrationsToMigrate(availableMigrations, migratedMigrations)

	basePath := []string{"db", "migrations"}

	for _, migration := range migrations {
		basePath := append(basePath, migration, "up.sql")

		if m.FileDirChecker.FileDirExists(strings.Join(basePath, "/")) == false {
			return errors.New(fmt.Sprintf("migrate: file %s does not exist", migration))
		}

		stmt, err := m.FileReader.Read(strings.Join(basePath, "/"))
		if err != nil {
			return errors.New(fmt.Sprintf("migrate: file %s could not be read: %v", migration, err))
		}

		_, err = m.DB.Exec(stmt)
		if err != nil {
			return errors.New(fmt.Sprintf("migrate: could not execute migration: %v", err))
		}

		timestamp := m.getTimestampFromMigrationName(migration)

		_, err = m.DB.Exec(fmt.Sprintf(migrationInsertStmt, timestamp))
		if err != nil {
			return errors.New(fmt.Sprintf("migrate: could not insert migration timestamp: %v", err))
		}
	}

	return nil
}

func (m *Migrator) fetchAvailableMigrations() ([]string, error) {
	if m.FileDirChecker.FileDirExists("db/migrations") == false {
		return nil, errors.New("migrate: migrations directory does not exist")
	}

	dirs, err := m.DirReader.Read("db/migrations")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("migrate: could not read migrations directory: %v", err))
	}

	names := make([]string, len(dirs))

	for i, v := range dirs {
		names[i] = v.Name()
	}

	return names, nil
}

func (m *Migrator) getMigrationsToMigrate(available, migrated []string) []string {
	var migrationsToMigrate []string

	for _, avail := range available {
		timestamp := m.getTimestampFromMigrationName(avail)

		if sliceContains(migrated, timestamp) == false {
			migrationsToMigrate = append(migrationsToMigrate, avail)
		}
	}

	return migrationsToMigrate
}

func (m *Migrator) getTimestampFromMigrationName(migration string) string {
	return strings.Split(migration, "_")[0]
}

func sliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}


package util

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const schemaMigrationsQuery = "SELECT migration FROM schema_migrations ;"

func InitDB() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("initdb: could not open connection to database: %v", err))
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("initdb: could not ping database: %v", err))
	}

	return db, nil
}

func GetSchemaMigrations(db *sql.DB) ([]string, error) {
	rows, err := db.Query(schemaMigrationsQuery)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("db: could not query schema_migrations: %v", err))
	}

	defer rows.Close()

	var migrations []string

	for rows.Next() {
		var migration string

		if err = rows.Scan(&migration); err != nil {
			return nil, errors.New(fmt.Sprintf("db: could not query schema_migrations: %v", err))
		}

		migrations = append(migrations, migration)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("db: could not query schema_migrations: %v", err))
	}

	return migrations, nil

}

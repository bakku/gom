package util

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

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

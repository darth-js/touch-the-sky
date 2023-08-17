package db

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

// Matches strings starting with '/', followed by alphanumeric characters, underscores, or slashes, and ending with '.db'.
const dbNameRegex = `^[/a-zA-Z0-9_]+\.db$`

func Connect(dbPath string) (*sql.DB, error) {

	if ok := regexp.MustCompile(dbNameRegex).MatchString(dbPath); !ok {
		return nil, fmt.Errorf("invalid db path: %s", dbPath)
	}

	var err error
	var db *sql.DB
	if db, err = sql.Open("sqlite3", dbPath); err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

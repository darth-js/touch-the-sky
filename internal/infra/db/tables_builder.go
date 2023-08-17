package db

import "database/sql"

type tablesBuilder struct {
	db         *sql.DB
	statements []string
}

func NewTablesBuilder(db *sql.DB) *tablesBuilder {
	return &tablesBuilder{db: db}
}

func (t *tablesBuilder) Build() error {
	for _, statement := range t.statements {
		_, err := t.db.Exec(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *tablesBuilder) WithUsersTable() *tablesBuilder {
	createStatement := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	t.statements = append(t.statements, createStatement)
	return t
}

func (t *tablesBuilder) WithVideosTable() *tablesBuilder {
	createStatement := `
	CREATE TABLE IF NOT EXISTS videos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		description TEXT,
		link TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`
	t.statements = append(t.statements, createStatement)
	return t
}

func (t *tablesBuilder) WithAnnotationsTable() *tablesBuilder {
	createStatement := `
	CREATE TABLE IF NOT EXISTS annotations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		video_id INTEGER NOT NULL,
		start_time TEXT NOT NULL,
		end_time TEXT NOT NULL,
		note TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE
	);
	`
	t.statements = append(t.statements, createStatement)
	return t
}

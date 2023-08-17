package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTablesBuilder_Build(t *testing.T) {
	// fixtures
	expectedTableNames := []string{"sqlite_sequence", "users", "videos", "annotations"}

	dbPath := TestDbPath
	defer Cleanup(dbPath)

	db := connectDb(dbPath, t)
	defer db.Close()

	builder := NewTablesBuilder(db)

	// test
	err := builder.WithUsersTable().WithVideosTable().WithAnnotationsTable().Build()

	// assert
	require.NoError(t, err)
	require.ElementsMatch(t, expectedTableNames, getDbTableNames(db, t))
}

func TestTablesBuilder_WithUsersTable(t *testing.T) {
	// fixtures
	builder := NewTablesBuilder(nil).WithUsersTable()

	// assert
	require.Len(t, builder.statements, 1)
	require.Contains(t, builder.statements[0], "CREATE TABLE IF NOT EXISTS users")
}

func TestTablesBuilder_WithVideosTable(t *testing.T) {
	// fixtures
	builder := NewTablesBuilder(nil)
	builder.WithVideosTable()

	// assert
	require.Len(t, builder.statements, 1)
	require.Contains(t, builder.statements[0], "CREATE TABLE IF NOT EXISTS videos")
}

func TestTablesBuilder_WithAnnotationsTable(t *testing.T) {
	// fixtures
	builder := NewTablesBuilder(nil)
	builder.WithAnnotationsTable()

	// assert
	require.Len(t, builder.statements, 1)
	require.Contains(t, builder.statements[0], "CREATE TABLE IF NOT EXISTS annotations")
}

func connectDb(dbPath string, t *testing.T) *sql.DB {
	db, err := Connect(dbPath)
	require.NoError(t, err)
	return db
}

func getDbTableNames(db *sql.DB, t *testing.T) []string {
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	require.NoError(t, err)
	defer rows.Close()

	tables := []string{}
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		require.NoError(t, err)
		tables = append(tables, tableName)
	}
	return tables
}

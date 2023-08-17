package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnect_success_happyPath(t *testing.T) {
	// fixtures
	dbPath := TestDbPath
	defer Cleanup(dbPath)

	// test
	db, err := Connect(dbPath)
	if db != nil {
		defer db.Close()
	}

	// assert
	require.NoError(t, err)
	require.NotNil(t, db)
}

func TestConnect_invalidDbPath(t *testing.T) {
	// fixtures
	dbPath := ""
	defer Cleanup(dbPath)

	// test
	db, err := Connect(dbPath)
	if db != nil {
		defer db.Close()
	}

	// assert
	require.Error(t, err)
	require.Nil(t, db)
}

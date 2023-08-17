package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
)

func beforeEach(t *testing.T) {
	var err error
	db, mock, err = sqlmock.New()
	require.NoError(t, err)
}

func afterEach() {
	db.Close()
}

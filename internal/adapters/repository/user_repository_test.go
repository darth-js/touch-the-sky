package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"

	"github.com/stretchr/testify/require"
)

func TestUserRepository_FindByUsername_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	user := &model.User{
		Username:  "johndoe",
		Password:  "password123",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
		ID:        1,
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "email", "created_at"}).
		AddRow(user.ID, user.Username, user.Password, user.Email, user.CreatedAt)

	mock.ExpectQuery("^SELECT \\* FROM users WHERE username = \\?$").
		WithArgs(user.Username).
		WillReturnRows(rows)

	userRepo := NewUserRepository(db)

	// test
	result, err := userRepo.FindByUsername(user.Username)

	// assert
	require.NoError(t, err)
	require.Equal(t, user, result)
}

func TestUserRepository_FindByUsername_UnhappyPath_UserNotFound(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	username := "johndoe"

	mock.ExpectQuery("^SELECT \\* FROM users WHERE username = \\?$").
		WithArgs(username).
		WillReturnError(sql.ErrNoRows)

	userRepo := NewUserRepository(db)

	// test
	result, err := userRepo.FindByUsername(username)

	// assert
	require.Nil(t, result)
	require.EqualError(t, err, UserNotFoundError.Error())
}

func TestUserRepository_FindByUsername_UnhappyPath_DatabaseError(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	username := "johndoe"

	mock.ExpectQuery("^SELECT \\* FROM users WHERE username = \\?$").
		WithArgs(username).
		WillReturnError(errors.New("database error"))

	userRepo := NewUserRepository(db)

	// test
	result, err := userRepo.FindByUsername(username)

	// assert
	require.Nil(t, result)
	require.EqualError(t, err, "database error")
}

func TestUserRepository_Save_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	user := &model.User{
		Username:  "johndoe",
		Password:  "password123",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
	}

	mock.ExpectExec("^INSERT INTO users \\(username, password, email, created_at\\) VALUES \\(\\?, \\?, \\?, \\?\\)$").
		WithArgs(user.Username, user.Password, user.Email, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	userRepo := NewUserRepository(db)

	// test
	err := userRepo.Save(user)

	// assert
	require.NoError(t, err)
}

func TestUserRepository_Save_UnhappyPath_DatabaseError(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	user := &model.User{
		Username:  "johndoe",
		Password:  "password123",
		Email:     "johndoe@example.com",
		CreatedAt: time.Now(),
	}

	mock.ExpectExec("^INSERT INTO users \\(username, password, email, created_at\\) VALUES \\(\\?, \\?, \\?, \\?\\)$").
		WithArgs(user.Username, user.Password, user.Email, user.CreatedAt).
		WillReturnError(errors.New("database error"))

	userRepo := NewUserRepository(db)

	// test
	err := userRepo.Save(user)

	// assert
	require.EqualError(t, err, "database error")
}

package service

import (
	"testing"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestComparePassword_HappyPath(t *testing.T) {
	// fixture
	hashedPassword := "$2a$10$7GRkdPm.s1IrBpXKlb.SOu7vOIvFKUG0H/QSJEGCZVzHqq/ZSbBW."
	rawPassword := "password123"

	// test
	err := comparePassword(hashedPassword, rawPassword)

	// assert
	require.NoError(t, err)
}

func TestComparePassword_UnhappyPath(t *testing.T) {
	// fixture
	hashedPassword := "$2a$10$7GRkdPm.s1IrBpXKlb.SOu7vOIvFKUG0H/QSJEGCZVzHqq/ZSbBW."
	rawPassword := "wrongpassword"
	expectedError := bcrypt.ErrMismatchedHashAndPassword.Error()

	// test
	err := comparePassword(hashedPassword, rawPassword)

	// assert
	require.EqualError(t, err, expectedError)
}

func TestExtractUserName_HappyPath(t *testing.T) {
	// fixture
	email := "johndoe@example.com"
	expectedUsername := "johndoe"

	// test
	username, err := extractUserName(email)

	// assert
	require.NoError(t, err)
	require.Equal(t, expectedUsername, username)
}

func TestExtractUserName_UnhappyPath(t *testing.T) {
	// fixture
	email := "johndoeexample.com"
	expectedError := validation.ErrEmailIsInvalid.Error()

	// test
	username, err := extractUserName(email)

	// assert
	require.EqualError(t, err, expectedError)
	require.Empty(t, username)
}

func TestHashPassword_HappyPath(t *testing.T) {
	// fixture
	rawPassword := "password123"

	// test
	hashedPassword, err := hashPassword(rawPassword)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
}

func TestHashPassword_UnhappyPath(t *testing.T) {
	// fixture
	rawPassword := ""
	expectedError := validation.ErrPasswordIsInvalid.Error()

	// test
	hashedPassword, err := hashPassword(rawPassword)

	// assert
	require.EqualError(t, err, expectedError)
	require.Empty(t, hashedPassword)
}

package validation

import (
	"testing"
	"time"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/stretchr/testify/require"
)

func TestValidateUser_HappyPath(t *testing.T) {
	// fixtures
	user := &model.User{
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
	}

	// test
	err := ValidateUser(user)

	// assert
	require.NoError(t, err)
}

func TestValidateUser_NilUser(t *testing.T) {
	// test
	err := ValidateUser(nil)

	// assert
	require.EqualError(t, err, ErrUserIsNil.Error())
}

func TestValidateUser_InvalidUsername(t *testing.T) {
	// fixtures
	user := &model.User{
		Username:  "",
		Email:     "johndoe@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
	}

	// test
	err := ValidateUser(user)

	// assert
	require.EqualError(t, err, ErrNameIsInvalid.Error())
}

func TestValidateUser_InvalidEmail(t *testing.T) {
	//fixtures
	user := &model.User{
		Username:  "johndoe",
		Email:     "johndoeexample.com",
		Password:  "password123",
		CreatedAt: time.Now(),
	}

	// test
	err := ValidateUser(user)

	// assert
	require.EqualError(t, err, ErrEmailIsInvalid.Error())
}

func TestValidateUser_InvalidPassword(t *testing.T) {
	//fixtures
	user := &model.User{
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "",
		CreatedAt: time.Now(),
	}

	// test
	err := ValidateUser(user)

	// assert
	require.EqualError(t, err, ErrPasswordIsInvalid.Error())
}

func TestValidateUser_InvalidCreatedAt(t *testing.T) {
	//fixtures
	user := &model.User{
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "password123",
		CreatedAt: time.Time{},
	}

	// test
	err := ValidateUser(user)

	// assert
	require.EqualError(t, err, ErrUserCreatedAtIsInvalid.Error())
}

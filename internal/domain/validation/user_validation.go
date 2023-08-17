package validation

import (
	"fmt"
	"regexp"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
)

var (
	ErrUserIsNil              = fmt.Errorf("user is nil")
	ErrNameIsInvalid          = fmt.Errorf("name is invalid")
	ErrEmailIsInvalid         = fmt.Errorf("email is invalid")
	ErrPasswordIsInvalid      = fmt.Errorf("password is invalid")
	ErrUserCreatedAtIsInvalid = fmt.Errorf("created_at is invalid")

	UserValidationErrors = map[error]bool{
		ErrUserIsNil:              true,
		ErrNameIsInvalid:          true,
		ErrEmailIsInvalid:         true,
		ErrPasswordIsInvalid:      true,
		ErrUserCreatedAtIsInvalid: true,
	}

	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func ValidateUser(user *model.User) error {
	if user == nil {
		return ErrUserIsNil
	}
	if user.Username == "" {
		return ErrNameIsInvalid
	}

	if !emailRegex.MatchString(user.Email) {
		return ErrEmailIsInvalid
	}

	if user.Password == "" {
		return ErrPasswordIsInvalid
	}

	if user.CreatedAt.IsZero() {
		return ErrUserCreatedAtIsInvalid
	}
	return nil
}

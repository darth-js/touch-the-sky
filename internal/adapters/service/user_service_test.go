package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
	"github.com/juliocnsouzadev/go-videos-api/internal/infra/auth"
	"github.com/stretchr/testify/require"
)

var (
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
)

func TestUserService_Login_HappyPath(t *testing.T) {
	// fixture
	userRepo := &mockUserRepository{
		users: map[string]*model.User{
			"johndoe": {
				Username: "johndoe",
				Password: "$2a$10$7GRkdPm.s1IrBpXKlb.SOu7vOIvFKUG0H/QSJEGCZVzHqq/ZSbBW.",
			},
		},
	}

	authService := auth.NewAuthService("secret-key")
	userService := NewUserService(userRepo, authService)

	// test
	token, err := userService.Login("johndoe", "password123")

	// assertions
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestUserService_Login_UnhappyPath_UserNotFound(t *testing.T) {
	// fixture
	userRepo := &mockUserRepository{
		users: map[string]*model.User{},
	}
	authService := auth.NewAuthService("secret-key")
	userService := NewUserService(userRepo, authService)

	// test
	token, err := userService.Login("johndoe", "password123")

	// assertions
	require.EqualError(t, err, UserOrPasswordNotFoundError.Error())
	require.Empty(t, token)
}

func TestUserService_Login_UnhappyPath_InvalidPassword(t *testing.T) {
	// fixture
	userRepo := &mockUserRepository{
		users: map[string]*model.User{
			"johndoe": {
				Username: "johndoe",
				Password: "$2a$10$7GRkdPm.s1IrBpXKlb.SOu7vOIvFKUG0H/QSJEGCZVzHqq/ZSbBW.",
			},
		},
	}
	authService := auth.NewAuthService("secret-key")
	userService := NewUserService(userRepo, authService)

	// test
	token, err := userService.Login("johndoe", "wrongpassword")

	// assertions
	require.EqualError(t, err, UserOrPasswordNotFoundError.Error())
	require.Empty(t, token)
}

func TestUserService_Signup_HappyPath(t *testing.T) {
	// fixture
	userRepo := &mockUserRepository{
		users: map[string]*model.User{},
	}
	authService := auth.NewAuthService("secret-key")
	userService := NewUserService(userRepo, authService)

	// test
	token, err := userService.Signup("johndoe@example.com", "password123")

	// assertions
	require.NoError(t, err)
	require.NotEmpty(t, token)

	user, err := userRepo.FindByUsername("johndoe")
	require.NoError(t, err)
	require.Equal(t, "johndoe", user.Username)
	require.NotEmpty(t, user.Password)
	require.WithinDuration(t, time.Now(), user.CreatedAt, 1*time.Second)
}

func TestUserService_Signup_UnhappyPath_UserAlreadyExists(t *testing.T) {
	// fixture
	userRepo := &mockUserRepository{
		users: map[string]*model.User{
			"johndoe": {
				Username: "johndoe",
				Password: "$2a$10$7GRkdPm.s1IrBpXKlb.SOu7vOIvFKUG0H/QSJEGCZVzHqq/ZSbBW.",
			},
		},
	}
	authService := auth.NewAuthService("secret-key")
	userService := NewUserService(userRepo, authService)

	// test
	token, err := userService.Signup("johndoe@example.com", "password123")

	// assertions
	require.EqualError(t, err, ErrUserAlreadyExists.Error())
	require.Empty(t, token)
}

func TestUserService_Signup_UnhappyPath_InvalidEmail(t *testing.T) {
	// fixture
	userRepo := &mockUserRepository{
		users: map[string]*model.User{},
	}
	authService := auth.NewAuthService("secret-key")
	userService := NewUserService(userRepo, authService)

	// test
	token, err := userService.Signup("johndoeexample.com", "password123")

	// assertions
	require.EqualError(t, err, validation.ErrEmailIsInvalid.Error())
	require.Empty(t, token)
}

func TestUserService_Signup_UnhappyPath_InvalidPassword(t *testing.T) {
	//fixture
	userRepo := &mockUserRepository{
		users: map[string]*model.User{},
	}
	authService := auth.NewAuthService("secret-key")
	userService := NewUserService(userRepo, authService)

	// test
	token, err := userService.Signup("johndoe@example.com", "")

	// assertions
	require.EqualError(t, err, validation.ErrPasswordIsInvalid.Error())
	require.Empty(t, token)
}

type mockUserRepository struct {
	users map[string]*model.User
}

func (r *mockUserRepository) FindByUsername(username string) (*model.User, error) {
	user, ok := r.users[username]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *mockUserRepository) Save(user *model.User) error {
	if _, ok := r.users[user.Username]; ok {
		return ErrUserAlreadyExists
	}
	r.users[user.Username] = user
	return nil
}

package service

import (
	"fmt"
	"log"
	"time"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/ports"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
	"github.com/juliocnsouzadev/go-videos-api/internal/infra/auth"
)

var UserOrPasswordNotFoundError = fmt.Errorf("invalid username or password")

type userService struct {
	userRepo ports.UserRepository
	auth     auth.AuthService
}

func NewUserService(userRepo ports.UserRepository, auth auth.AuthService) *userService {
	return &userService{
		userRepo: userRepo,
		auth:     auth,
	}
}

func (s *userService) Login(username string, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", UserOrPasswordNotFoundError
	}

	if err := comparePassword(user.Password, password); err != nil {
		log.Printf("error comparing password: %v", err)
		return "", UserOrPasswordNotFoundError
	}

	return s.createSession(user)
}

func (s *userService) Signup(email string, password string) (string, error) {
	var err error

	user := &model.User{
		CreatedAt: time.Now(),
		Email:     email,
	}

	if user.Password, err = hashPassword(password); err != nil {
		return "", err
	}
	if user.Username, err = extractUserName(email); err != nil {
		return "", err
	}

	if err = validation.ValidateUser(user); err != nil {
		return "", err
	}

	if err = s.userRepo.Save(user); err != nil {
		return "", err
	}

	return s.createSession(user)
}

func (s *userService) createSession(user *model.User) (string, error) {
	return s.auth.GenerateJwtToken(user.Username)
}

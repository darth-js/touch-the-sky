package service

import (
	"strings"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
	"golang.org/x/crypto/bcrypt"
)

func comparePassword(hashedPasword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", validation.ErrPasswordIsInvalid
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func extractUserName(email string) (string, error) {
	tokens := strings.Split(email, "@")
	if (tokens == nil) || (len(tokens) < 2) {
		return "", validation.ErrEmailIsInvalid
	}
	return tokens[0], nil
}

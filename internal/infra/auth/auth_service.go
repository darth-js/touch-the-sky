package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthService interface {
	GenerateJwtToken(username string) (string, error)
	ValidateJwtToken(tokenString string) (bool, string)
}

type authService struct {
	jwtKey []byte
}

func NewAuthService(key string) *authService {
	return &authService{
		jwtKey: []byte(key),
	}
}

func (a *authService) GenerateJwtToken(username string) (string, error) {
	// Set the JWT claims
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authService) ValidateJwtToken(tokenString string) (bool, string) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return a.jwtKey, nil
	})

	if err != nil {
		return false, ""
	}

	return token.Valid, claims.Username
}

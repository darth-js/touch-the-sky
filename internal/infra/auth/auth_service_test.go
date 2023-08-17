package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestAuthService_GenerateJwtToken_HappyPath(t *testing.T) {
	// fixture
	authService := NewAuthService("secret-key")
	username := "johndoe"

	// test
	token, err := authService.GenerateJwtToken(username)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestAuthService_ValidateJwtToken_HappyPath(t *testing.T) {
	// fixture
	authService := NewAuthService("secret-key")
	token := generateTestToken(t, authService)

	// test
	valid, _ := authService.ValidateJwtToken(token)

	// assert
	require.True(t, valid)
}

func generateTestToken(t *testing.T, authService AuthService) string {
	username := "johndoe"
	token, err := authService.GenerateJwtToken(username)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	return token
}

func TestAuthService_ValidateJwtToken_UnhappyPath_InvalidToken(t *testing.T) {
	// fixture
	authService := NewAuthService("secret-key")

	token := "invalid-token"

	// test
	valid, _ := authService.ValidateJwtToken(token)

	// assert
	require.False(t, valid)
}

func TestAuthService_ValidateJwtToken_UnhappyPath_ExpiredToken(t *testing.T) {
	// fixture
	authService := NewAuthService("secret-key")

	expiredToke := generateExpiredTestToken(t, authService)

	// test
	valid, _ := authService.ValidateJwtToken(expiredToke)

	// assert
	require.False(t, valid)
}

func generateExpiredTestToken(t *testing.T, service AuthService) string {
	// Set the expiration time to 1 second ago
	expirationTime := time.Now().Add(-1 * time.Second)

	claims := &Claims{
		Username: "johndoe",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	authServiceImpl := service.(*authService)

	tokenString, err := token.SignedString(authServiceImpl.jwtKey)
	require.NoError(t, err)

	return tokenString
}

func TestAuthService_ValidateJwtToken_UnhappyPath_InvalidKey(t *testing.T) {
	// fixture
	authService1 := NewAuthService("secret-key-1")
	authService2 := NewAuthService("secret-key-2")

	token := generateTestToken(t, authService1)

	valid, _ := authService2.ValidateJwtToken(token)
	require.False(t, valid)
}

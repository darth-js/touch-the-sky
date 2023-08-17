package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/juliocnsouzadev/go-videos-api/internal/adapters/service"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_SignupHandler_HappyPath(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	payload := map[string]string{
		"email":    "new-user@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(body))
	respWriter := httptest.NewRecorder()

	// test
	handler.SignupHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusOK, respWriter.Code)

	var response map[string]string
	_ = json.Unmarshal(respWriter.Body.Bytes(), &response)

	require.Equal(t, "token", response["token"])
}

func TestUserHandler_SignupHandler_UnhappyPath_MethodNotAllowed(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	req, _ := http.NewRequest("GET", "/signup", nil)
	respWriter := httptest.NewRecorder()

	// test
	handler.SignupHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusMethodNotAllowed, respWriter.Code)
}

func TestUserHandler_SignupHandler_UnhappyPath_InvalidRequestPayload(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader([]byte("invalid-payload")))
	respWriter := httptest.NewRecorder()

	// test
	handler.SignupHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusBadRequest, respWriter.Code)
}

func TestUserHandler_SignupHandler_UnhappyPath_UserOrPasswordNotFoundError(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	payload := map[string]string{
		"email":    "existing-user@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(body))
	respWriter := httptest.NewRecorder()

	// test
	handler.SignupHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusUnauthorized, respWriter.Code)
}

func TestUserHandler_SignupHandler_UnhappyPath_ValidationError(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	payload := map[string]string{
		"email":    "invalid-email",
		"password": "password",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(body))
	respWriter := httptest.NewRecorder()

	// test
	handler.SignupHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusBadRequest, respWriter.Code)
	require.Contains(t, respWriter.Body.String(), validation.ErrEmailIsInvalid.Error())
}

func TestUserHandler_LoginHandler_HappyPath(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	payload := map[string]string{
		"email":    "existing-user@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
	respWriter := httptest.NewRecorder()

	// test
	handler.LoginHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusOK, respWriter.Code)

	var response map[string]string
	_ = json.Unmarshal(respWriter.Body.Bytes(), &response)

	require.Equal(t, "token", response["token"])
}

func TestUserHandler_LoginHandler_UnhappyPath_MethodNotAllowed(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	req, _ := http.NewRequest("GET", "/login", nil)
	respWriter := httptest.NewRecorder()

	// test
	handler.LoginHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusMethodNotAllowed, respWriter.Code)
}

func TestUserHandler_LoginHandler_UnhappyPath_InvalidRequestPayload(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	req, _ := http.NewRequest("POST", "/login", bytes.NewReader([]byte("invalid-payload")))
	respWriter := httptest.NewRecorder()

	// test
	handler.LoginHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusBadRequest, respWriter.Code)
}

func TestUserHandler_LoginHandler_UnhappyPath_UserOrPasswordNotFoundError(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	payload := map[string]string{
		"email":    "non-existing-user@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
	respWriter := httptest.NewRecorder()

	// test
	handler.LoginHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusUnauthorized, respWriter.Code)
}

func TestUserHandler_LoginHandler_UnhappyPath_InternalServerError(t *testing.T) {
	// fixture
	handler := NewUserHandler(&mockUserService{})

	payload := map[string]string{
		"email":    "existing-user@example.com",
		"password": "invalid-password",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
	respWriter := httptest.NewRecorder()

	// test
	handler.LoginHandler(respWriter, req)

	// assertion
	require.Equal(t, http.StatusUnauthorized, respWriter.Code)
}

type mockUserService struct{}

func (s *mockUserService) Signup(email, password string) (string, error) {
	if email == "existing-user@example.com" {
		return "", service.UserOrPasswordNotFoundError
	}
	if email == "invalid-email" {
		return "", validation.ErrEmailIsInvalid
	}
	return "token", nil
}

func (s *mockUserService) Login(email, password string) (string, error) {
	if email == "non-existing-user@example.com" || password == "invalid-password" {
		return "", service.UserOrPasswordNotFoundError
	}

	return "token", nil
}

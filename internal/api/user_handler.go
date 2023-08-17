package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juliocnsouzadev/go-videos-api/internal/adapters/service"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/ports"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(us ports.UserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (h *UserHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userDto := &UserDto{}

	if err := json.NewDecoder(r.Body).Decode(userDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var token string
	var err error
	if token, err = h.userService.Signup(userDto.Email, userDto.Password); err != nil {
		statusCode := http.StatusInternalServerError
		if err == service.UserOrPasswordNotFoundError {
			statusCode = http.StatusUnauthorized
		}

		if _, ok := validation.UserValidationErrors[err]; ok {
			statusCode = http.StatusBadRequest
			msg := fmt.Sprintf("Singnup failed due %s", err.Error())
			http.Error(w, msg, statusCode)
			return
		}

		http.Error(w, "Signup failed", statusCode)
		return
	}

	respondWithToken(w, token)
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userDto := &UserDto{}
	if err := json.NewDecoder(r.Body).Decode(userDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var token string
	var err error
	if token, err = h.userService.Login(userDto.Email, userDto.Password); err != nil {
		statusCode := http.StatusInternalServerError
		if err == service.UserOrPasswordNotFoundError {
			statusCode = http.StatusUnauthorized
		}
		http.Error(w, "Login failed", statusCode)
		return
	}

	respondWithToken(w, token)
}

func respondWithToken(w http.ResponseWriter, token string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juliocnsouzadev/go-videos-api/internal/adapters/service"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/ports"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/validation"
	"github.com/juliocnsouzadev/go-videos-api/internal/infra/auth"
)

type VideoHandler struct {
	videoService ports.VideoService
	authService  auth.AuthService
}

func NewVideoHandler(service ports.VideoService, authService auth.AuthService) *VideoHandler {
	return &VideoHandler{
		videoService: service,
		authService:  authService,
	}
}

func (h *VideoHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")

	var username string
	var ok bool
	if ok, username = h.authService.ValidateJwtToken(tokenString); !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	videoDto := &VideoDto{}

	if err := json.NewDecoder(r.Body).Decode(videoDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.videoService.Create(username, &videoDto.Video, videoDto.Annotaions); err != nil {
		statusCode := http.StatusInternalServerError

		if _, ok := validation.VideoValidationErrors[err]; !ok {
			statusCode = http.StatusBadRequest
			msg := fmt.Sprintf("Request failed due %s", err.Error())
			http.Error(w, msg, statusCode)
			return
		}

		if _, ok := validation.AnnotationValidationErrors[err]; !ok {
			statusCode = http.StatusBadRequest
			msg := fmt.Sprintf("Request failed due %s", err.Error())
			http.Error(w, msg, statusCode)
			return
		}

		http.Error(w, "Request failed", statusCode)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *VideoHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	if ok, _ := h.authService.ValidateJwtToken(tokenString); !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	videoIdStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	videoDto := &VideoDto{}

	if err := json.NewDecoder(r.Body).Decode(videoDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.videoService.Update(int(videoId), &videoDto.Video, videoDto.Annotaions); err != nil {
		statusCode := http.StatusInternalServerError

		if _, ok := validation.VideoValidationErrors[err]; ok {
			statusCode = http.StatusBadRequest
			msg := fmt.Sprintf("Request failed due %s", err.Error())
			http.Error(w, msg, statusCode)
			return
		}

		if _, ok := validation.AnnotationValidationErrors[err]; ok {
			statusCode = http.StatusBadRequest
			msg := fmt.Sprintf("Request failed due %s", err.Error())
			http.Error(w, msg, statusCode)
			return
		}

		http.Error(w, "Request failed", statusCode)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *VideoHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	if ok, _ := h.authService.ValidateJwtToken(tokenString); !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	videoIdStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	video, annotations, err := h.videoService.Find(videoId)
	if err != nil {
		if err == service.ErrVideoNotFound || err == service.ErrAnnotationsNotFound {
			http.Error(w, "Video not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Request failed", http.StatusInternalServerError)
		return
	}

	videoDto := &VideoDto{
		Video:      *video,
		Annotaions: annotations,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videoDto)

}

func (h *VideoHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	if ok, _ := h.authService.ValidateJwtToken(tokenString); !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	videoIdStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.videoService.Remove(int(videoId)); err != nil {
		http.Error(w, "Request failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}

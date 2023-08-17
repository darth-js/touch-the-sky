package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/ports"
	"github.com/juliocnsouzadev/go-videos-api/internal/infra/auth"
)

func StartHttpServer(
	authService auth.AuthService,
	userService ports.UserService,
	videoService ports.VideoService) {
	router := mux.NewRouter()
	userHandler := NewUserHandler(userService)

	router.HandleFunc("/signup", userHandler.SignupHandler).Methods("POST")
	router.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")

	videorHandler := NewVideoHandler(videoService, authService)
	router.HandleFunc("/videos/", videorHandler.CreateHandler).Methods("POST")
	router.HandleFunc("/videos/{id}/", videorHandler.UpdateHandler).Methods("PUT")
	router.HandleFunc("/videos/{id}/", videorHandler.GetHandler).Methods("GET")
	router.HandleFunc("/videos/{id}/", videorHandler.DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", nil)
}

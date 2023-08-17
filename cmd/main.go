package main

import (
	"database/sql"
	"log"

	"github.com/juliocnsouzadev/go-videos-api/internal/adapters/repository"
	"github.com/juliocnsouzadev/go-videos-api/internal/adapters/service"
	"github.com/juliocnsouzadev/go-videos-api/internal/api"
	"github.com/juliocnsouzadev/go-videos-api/internal/infra/auth"
	"github.com/juliocnsouzadev/go-videos-api/internal/infra/config"
	"github.com/juliocnsouzadev/go-videos-api/internal/infra/db"
)

var database *sql.DB

func main() {
	defer Cleanup()

	log.Println("Starting server...")

	var err error
	var settings *config.Settings

	log.Println("Loading settings...")
	if settings, err = config.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to database...")
	if database, err = db.Connect(settings.DatabaseURL); err != nil {
		log.Fatal(err)
	}

	log.Println("Creating tables...")
	if err = createTables(database); err != nil {
		log.Fatal(err)
	}

	authService := auth.NewAuthService(settings.JwtKey)
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository, authService)

	videoRepo := repository.NewVideoRepository(database)
	annotationRepo := repository.NewAnnotationRepository(database)
	videoService := service.NewVideoService(videoRepo, annotationRepo, userRepository)

	log.Println("Starting HTTP server...")
	api.StartHttpServer(authService, userService, videoService)

	log.Println("Server started")
}

func createTables(database *sql.DB) error {
	return db.
		NewTablesBuilder(database).
		WithUsersTable().
		WithVideosTable().
		WithAnnotationsTable().
		Build()
}

func Cleanup() {
	log.Println("Closing database connection...")
	database.Close()
}

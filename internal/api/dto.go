package api

import "github.com/juliocnsouzadev/go-videos-api/internal/domain/model"

type UserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VideoDto struct {
	model.Video
	Annotaions []*model.Annotation
}

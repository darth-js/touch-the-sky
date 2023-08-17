package ports

import "github.com/juliocnsouzadev/go-videos-api/internal/domain/model"

type VideoService interface {
	Create(username string, video *model.Video, annotaions []*model.Annotation) error
	Find(int) (*model.Video, []*model.Annotation, error)
	Update(int, *model.Video, []*model.Annotation) error
	Remove(int) error
}

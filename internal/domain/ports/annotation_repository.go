package ports

import (
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
)

type AnnotationRepository interface {
	Create(annotation *model.Annotation, videoId, userId int) error
	FindVideoId(int) ([]*model.Annotation, error)
	Update(int, *model.Annotation) error
	Remove(int) error
}

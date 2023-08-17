package ports

import "github.com/juliocnsouzadev/go-videos-api/internal/domain/model"

type VideoRepository interface {
	Create(video *model.Video, userId int) (int, error)
	FindById(int) (*model.Video, error)
	Update(int, *model.Video) error
	Remove(int) error
}

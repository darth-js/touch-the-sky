package ports

import "github.com/juliocnsouzadev/go-videos-api/internal/domain/model"

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	Save(user *model.User) error
}

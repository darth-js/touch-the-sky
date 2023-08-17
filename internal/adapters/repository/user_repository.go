package repository

import (
	"database/sql"
	"fmt"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
)

var (
	UserNotFoundError = fmt.Errorf("user not found")
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT * FROM users WHERE username = ?`
	err := u.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, UserNotFoundError
		}
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Save(user *model.User) error {
	query := `INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?)`
	_, err := u.db.Exec(query, user.Username, user.Password, user.Email, user.CreatedAt)
	return err
}

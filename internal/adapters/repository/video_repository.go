package repository

import (
	"database/sql"
	"fmt"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
)

var (
	VideoNotFoundError = fmt.Errorf("video not found")
)

type videoRepository struct {
	db *sql.DB
}

func NewVideoRepository(db *sql.DB) *videoRepository {
	return &videoRepository{db}
}

func (r *videoRepository) Create(video *model.Video, userId int) (int, error) {
	query := `INSERT INTO videos (title, description, description, user_id, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, video.Title, video.Description, video.Link, userId, video.CreatedAt)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), nil
}

func (r *videoRepository) FindById(id int) (*model.Video, error) {

	video := &model.Video{}
	query := `SELECT * FROM videos WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&video.ID, &video.CreatedAt, &video.Duration,
		&video.Description, &video.Link, &video.Title, &video.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, VideoNotFoundError
		}
		return nil, err
	}
	return video, nil
}

func (r *videoRepository) Update(id int, video *model.Video) error {

	query := `UPDATE videos SET title = ?, description = ?, link = ? WHERE id = ?`
	_, err := r.db.Exec(query, video.Title, video.Description, video.Link, id)
	return err
}

func (r *videoRepository) Remove(id int) error {

	query := `DELETE FROM videos WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

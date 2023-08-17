package repository

import (
	"database/sql"
	"fmt"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
)

var (
	ErrAnnotationNotFound = fmt.Errorf("annotation not found")
)

type annotationRepository struct {
	db *sql.DB
}

func NewAnnotationRepository(db *sql.DB) *annotationRepository {
	return &annotationRepository{db}
}

func (r *annotationRepository) Create(annotation *model.Annotation, userId, videoId int) error {
	query := `INSERT INTO annotations (start_time, end_time, type, note, user_id, video_id) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, annotation.StartTime, annotation.EndTime, annotation.Type, annotation.Note, userId, videoId)
	return err
}

func (r *annotationRepository) FindVideoId(id int) ([]*model.Annotation, error) {

	annotations := []*model.Annotation{}
	query := `SELECT * FROM annotations WHERE id = ?`

	rows, err := r.db.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAnnotationNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		annotation := &model.Annotation{}
		err := rows.Scan(&annotation.ID, &annotation.StartTime, &annotation.EndTime,
			&annotation.Type, &annotation.Note, &annotation.UserID, &annotation.VideoID)
		if err != nil {
			return nil, err
		}
		annotations = append(annotations, annotation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return annotations, nil
}

func (r *annotationRepository) Update(id int, annotation *model.Annotation) error {

	query := `UPDATE annotations SET start_time = ?, end_time = ?, type = ?, note = ? WHERE id = ?`
	_, err := r.db.Exec(query, annotation.StartTime, annotation.EndTime, annotation.Type, annotation.Note, id)
	return err
}

func (r *annotationRepository) Remove(id int) error {

	query := `DELETE FROM annotations WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

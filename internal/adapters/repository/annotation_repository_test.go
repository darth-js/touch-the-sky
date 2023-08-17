package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/stretchr/testify/require"
)

func TestAnnotationRepository_Create_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	repo := NewAnnotationRepository(db)

	startTime := time.Duration(0)
	endTime := time.Duration(2)
	tp := "test"
	note := "test note"
	annotation := &model.Annotation{
		StartTime: startTime,
		EndTime:   endTime,
		Type:      tp,
		Note:      note,
	}
	videoId := 1
	userId := 1

	mock.ExpectExec("INSERT INTO annotations").
		WithArgs(startTime, endTime, tp, note, userId, videoId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// test
	err := repo.Create(annotation, userId, videoId)

	// assertions
	require.NoError(t, err)
}

func TestAnnotationRepository_Find_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	repo := NewAnnotationRepository(db)

	id := 1
	videoId := 1
	userId := 1
	startTime := time.Duration(0)
	endTime := time.Duration(2)
	tp := "test"
	note := "test note"

	rows := sqlmock.NewRows([]string{"id", "start_time", "end_time", "type", "note", "user_id", "video_id"}).
		AddRow(id, startTime, endTime, tp, note, userId, videoId)
	mock.ExpectQuery("SELECT \\* FROM annotations").
		WithArgs(id).
		WillReturnRows(rows)

	expected := []*model.Annotation{
		{
			ID:        id,
			VideoID:   videoId,
			UserID:    userId,
			StartTime: startTime,
			EndTime:   endTime,
			Type:      tp,
			Note:      note,
		},
	}

	// test
	annotation, err := repo.FindVideoId(id)

	// assertions
	require.NoError(t, err)
	require.Equal(t, expected, annotation)
}

func TestAnnotationRepository_Find_UnhappyPath_NotFound(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	repo := NewAnnotationRepository(db)

	id := 1

	mock.ExpectQuery("SELECT \\* FROM annotations").
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	// test
	annotation, err := repo.FindVideoId(id)

	// assertions
	require.EqualError(t, err, ErrAnnotationNotFound.Error())
	require.Nil(t, annotation)
}

func TestAnnotationRepository_Update_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	repo := NewAnnotationRepository(db)

	id := 1
	startTime := time.Duration(0)
	endTime := time.Duration(2)
	tp := "test"
	note := "test note"

	annotation := &model.Annotation{
		StartTime: startTime,
		EndTime:   endTime,
		Type:      tp,
		Note:      note,
	}

	mock.ExpectExec("UPDATE annotations").
		WithArgs(startTime, endTime, tp, note, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// test
	err := repo.Update(id, annotation)

	// assertions
	require.NoError(t, err)
}

func TestAnnotationRepository_Update_UnhappyPath_NotFound(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	repo := NewAnnotationRepository(db)

	id := 1
	startTime := time.Duration(0)
	endTime := time.Duration(2)
	tp := "test"
	note := "test note"

	annotation := &model.Annotation{
		StartTime: startTime,
		EndTime:   endTime,
		Type:      tp,
		Note:      note,
	}

	mock.ExpectExec("UPDATE annotations").
		WithArgs(startTime, endTime, tp, note, id).
		WillReturnError(sql.ErrNoRows)

	// test
	err := repo.Update(id, annotation)

	// assertions
	require.Error(t, err)
}

func TestAnnotationRepository_Remove_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	repo := NewAnnotationRepository(db)

	id := 1

	mock.ExpectExec("DELETE FROM annotations").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// test
	err := repo.Remove(id)

	// assertions
	require.NoError(t, err)
}

func TestAnnotationRepository_Remove_UnhappyPath_NotFound(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	repo := NewAnnotationRepository(db)

	id := 1

	mock.ExpectExec("DELETE FROM annotations").
		WithArgs(id).
		WillReturnError(errors.New("database error"))

	// test
	err := repo.Remove(id)

	// assertions
	require.Error(t, err)
}

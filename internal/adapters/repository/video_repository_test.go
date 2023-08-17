package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"

	"github.com/stretchr/testify/require"
)

func TestVideoRepository_Create_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	videoRepo := NewVideoRepository(db)

	video := &model.Video{
		Title:       "Test Video",
		Description: "This is a test video",
		Link:        "https://example.com/test.mp4",
		CreatedAt:   time.Now(),
	}
	userId := 1

	mock.ExpectExec("INSERT INTO videos").
		WithArgs(video.Title, video.Description, video.Link, userId, video.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// test
	_, err := videoRepo.Create(video, userId)

	// assertions
	require.NoError(t, err)
}

func TestVideoRepository_Create_UnhappyPath(t *testing.T) {
	// Unhappy path test: database error
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	videoRepo := NewVideoRepository(db)

	video := &model.Video{
		Title:       "Test Video",
		Description: "This is a test video",
		Link:        "https://example.com/test.mp4",
		CreatedAt:   time.Now(),
	}
	userId := 1

	mock.ExpectExec("INSERT INTO videos").WithArgs(video.Title, video.Description, video.Link, userId, video.CreatedAt).WillReturnError(errors.New("database error"))

	_, err = videoRepo.Create(video, userId)
	require.Error(t, err)
}

func TestVideoRepository_Find_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	videoRepo := NewVideoRepository(db)

	videoID := 1
	userID := 1
	video := &model.Video{
		Title:       "Test Video",
		Description: "This is a test video",
		Link:        "https://example.com/test.mp4",
		CreatedAt:   time.Now(),
		Duration:    time.Duration(1),
		ID:          videoID,
		UserID:      userID,
	}

	rows := sqlmock.
		NewRows([]string{"id", "created_at", "duration", "description", "link", "title", "user_id"}).
		AddRow(video.ID, video.CreatedAt, video.Duration, video.Description, video.Link, video.Title, video.UserID)

	mock.ExpectQuery("SELECT \\* FROM videos").WithArgs(videoID).WillReturnRows(rows)

	// test
	result, err := videoRepo.FindById(videoID)

	// assertions
	require.NoError(t, err)
	require.Equal(t, video, result)
}

func TestVideoRepository_Find_UnhappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	videoRepo := NewVideoRepository(db)

	videoID := 1

	mock.ExpectQuery("SELECT \\* FROM videos").WithArgs(videoID).WillReturnError(VideoNotFoundError)

	// test
	_, err := videoRepo.FindById(videoID)

	// assertions
	require.Error(t, err)
}

func TestVideoRepository_Update_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	videoRepo := NewVideoRepository(db)

	videoID := 1
	video := &model.Video{
		Title:       "Test Video",
		Description: "This is a test video",
		Link:        "https://example.com/test.mp4",
		CreatedAt:   time.Now(),
	}

	mock.ExpectExec("UPDATE videos").
		WithArgs(video.Title, video.Description, video.Link, videoID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// test
	err := videoRepo.Update(videoID, video)

	// assertions
	require.NoError(t, err)
}

func TestVideoRepository_Update_UnhappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	videoRepo := NewVideoRepository(db)

	videoID := 1
	video := &model.Video{
		Title:       "Test Video",
		Description: "This is a test video",
		Link:        "https://example.com/test.mp4",
		CreatedAt:   time.Now(),
	}

	mock.ExpectExec("UPDATE videos").
		WithArgs(video.Title, video.Description, video.Link, videoID).
		WillReturnError(errors.New("database error"))

	// test
	err := videoRepo.Update(videoID, video)

	// assertions
	require.Error(t, err)
}

func TestVideoRepository_Remove_HappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	videoRepo := NewVideoRepository(db)

	videoID := 1

	mock.ExpectExec("DELETE FROM videos").WithArgs(videoID).WillReturnResult(sqlmock.NewResult(0, 1))

	// test
	err := videoRepo.Remove(videoID)

	// assertions
	require.NoError(t, err)
}

func TestVideoRepository_Remove_UnhappyPath(t *testing.T) {
	beforeEach(t)
	defer afterEach()

	// fixture
	videoRepo := NewVideoRepository(db)

	videoID := 1

	mock.ExpectExec("DELETE FROM videos").WithArgs(videoID).WillReturnError(errors.New("database error"))

	// test
	err := videoRepo.Remove(videoID)

	// assertions
	require.Error(t, err)
}

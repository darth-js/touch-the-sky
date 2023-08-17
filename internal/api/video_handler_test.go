package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"

	"github.com/juliocnsouzadev/go-videos-api/internal/adapters/service"
	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestVideoHandler_CreateHandler(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	video := model.Video{
		Title:       "Test Video",
		Description: "Test Description",
		Link:        "https://example.com/test.mp4",
	}

	annotations := []*model.Annotation{
		{
			Note:      "Test Annotation 1",
			StartTime: 10,
			ID:        1,
			VideoID:   1,
			UserID:    1,
			EndTime:   20,
			Type:      "test-type",
		},
		{
			Note:      "Test Annotation 2",
			StartTime: 21,
			ID:        2,
			VideoID:   1,
			UserID:    1,
			EndTime:   25,
			Type:      "test-type",
		},
	}

	videoDto := VideoDto{
		Video:      video,
		Annotaions: annotations,
	}

	videoJson, _ := json.Marshal(videoDto)

	req, err := http.NewRequest("POST", "/videos", bytes.NewBuffer(videoJson))
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)

	videoServiceMock.On("Create", "test-user", &video, annotations).Return(nil)

	// Execute
	rr := httptest.NewRecorder()
	handler.CreateHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusCreated, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_UpdateHandler(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	video := model.Video{
		Title:       "Test Video",
		Description: "Test Description",
		Link:        "https://example.com/test.mp4",
	}

	annotations := []*model.Annotation{
		{
			Note:      "Test Annotation 1",
			StartTime: 10,
			ID:        1,
			VideoID:   1,
			UserID:    1,
			EndTime:   20,
			Type:      "test-type",
		},
		{
			Note:      "Test Annotation 2",
			StartTime: 21,
			ID:        2,
			VideoID:   1,
			UserID:    1,
			EndTime:   25,
			Type:      "test-type",
		},
	}

	videoDto := VideoDto{
		Video:      video,
		Annotaions: annotations,
	}

	videoJson, _ := json.Marshal(videoDto)

	req, err := http.NewRequest("PUT", "/videos/1", bytes.NewBuffer(videoJson))
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	videoServiceMock.On("Update", 1, &video, annotations).Return(nil)

	// Execute
	rr := httptest.NewRecorder()
	handler.UpdateHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusOK, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_GetHandler(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	video := model.Video{
		ID:          1,
		Title:       "Test Video",
		Description: "Test Description",
		Link:        "https://example.com/test.mp4",
		UserID:      1,
		Duration:    100,
		CreatedAt:   time.Now(),
	}

	annotations := []*model.Annotation{
		{
			Note:      "Test Annotation 1",
			StartTime: 10,
			ID:        1,
			VideoID:   1,
			UserID:    1,
			EndTime:   20,
			Type:      "test-type",
		},
		{
			Note:      "Test Annotation 2",
			StartTime: 21,
			ID:        2,
			VideoID:   1,
			UserID:    1,
			EndTime:   25,
			Type:      "test-type",
		},
	}

	videoServiceMock.On("Find", 1).Return(&video, annotations, nil)

	req, err := http.NewRequest("GET", "/videos/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	// Execute
	rr := httptest.NewRecorder()
	handler.GetHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusOK, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_DeleteHandler(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	req, err := http.NewRequest("DELETE", "/videos/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	videoServiceMock.On("Remove", 1).Return(nil)

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	// Execute
	rr := httptest.NewRecorder()
	handler.DeleteHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusAccepted, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_CreateHandler_InvalidRequestPayload(t *testing.T) {

	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	req, err := http.NewRequest("POST", "/videos", bytes.NewBufferString("invalid-json"))
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	// Execute
	rr := httptest.NewRecorder()
	handler.CreateHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_UpdateHandler_InvalidRequestPayload(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	req, err := http.NewRequest("PUT", "/videos/1", bytes.NewBufferString("invalid-json"))
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	// Execute
	rr := httptest.NewRecorder()
	handler.UpdateHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_UpdateHandler_InvalidVideoId(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	req, err := http.NewRequest("PUT", "/videos/invalid-id", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid-id"})

	// Execute
	rr := httptest.NewRecorder()
	handler.UpdateHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_GetHandler_InvalidVideoId(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	req, err := http.NewRequest("GET", "/videos/invalid-id", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid-id"})

	// Execute
	rr := httptest.NewRecorder()
	handler.GetHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_GetHandler_VideoNotFound(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	videoServiceMock.On("Find", 1).Return(nil, nil, service.ErrVideoNotFound)

	req, err := http.NewRequest("GET", "/videos/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	// Execute
	rr := httptest.NewRecorder()
	handler.GetHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusNotFound, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_GetHandler_AnnotationsNotFound(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	video := model.Video{
		ID:          1,
		Title:       "Test Video",
		Description: "Test Description",
		Link:        "https://example.com/test.mp4",
		UserID:      1,
		Duration:    100,
		CreatedAt:   time.Now(),
	}

	videoServiceMock.On("Find", 1).Return(&video, nil, service.ErrAnnotationsNotFound)

	req, err := http.NewRequest("GET", "/videos/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	// Execute
	rr := httptest.NewRecorder()
	handler.GetHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusNotFound, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

func TestVideoHandler_DeleteHandler_InvalidVideoId(t *testing.T) {
	// Setup
	videoServiceMock := new(VideoServiceMock)
	authServiceMock := new(AuthService)

	handler := NewVideoHandler(videoServiceMock, authServiceMock)

	req, err := http.NewRequest("DELETE", "/videos/invalid-id", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "test-token"
	authServiceMock.On("ValidateJwtToken", token).Return(true, "test-user")

	req.Header.Add("Authorization", token)
	req = mux.SetURLVars(req, map[string]string{"id": "invlaid-id"})

	// Execute
	rr := httptest.NewRecorder()
	handler.DeleteHandler(rr, req)

	// Verify
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	videoServiceMock.AssertExpectations(t)
	authServiceMock.AssertExpectations(t)
}

type VideoServiceMock struct {
	mock.Mock
}

func (s *VideoServiceMock) Create(username string, video *model.Video, annotations []*model.Annotation) error {
	args := s.Called(username, video, annotations)
	return args.Error(0)
}
func (s *VideoServiceMock) Find(id int) (*model.Video, []*model.Annotation, error) {
	args := s.Called(id)
	get0 := args.Get(0)
	get1 := args.Get(1)
	if get0 == nil || get1 == nil {
		return nil, nil, args.Error(2)
	}
	v := get0.(*model.Video)
	a := get1.([]*model.Annotation)
	return v, a, args.Error(2)
}
func (s *VideoServiceMock) Update(id int, video *model.Video, annotations []*model.Annotation) error {
	args := s.Called(id, video, annotations)
	return args.Error(0)
}
func (s *VideoServiceMock) Remove(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

type AuthService struct {
	mock.Mock
}

func (s *AuthService) GenerateJwtToken(username string) (string, error) {
	args := s.Called(username)
	return args.String(0), args.Error(1)
}

func (s *AuthService) ValidateJwtToken(tokenString string) (bool, string) {
	args := s.Called(tokenString)
	return args.Bool(0), args.String(1)
}

package validation

import (
	"testing"
	"time"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/stretchr/testify/require"
)

func TestValidateVideo_HappyPath(t *testing.T) {
	// fixture
	video := &model.Video{
		Title:       "test title",
		Description: "test description",
		Link:        "https://example.com/test",
		UserID:      1,
		Duration:    time.Duration(1),
		CreatedAt:   time.Now(),
	}

	// test
	err := ValidateVideo(video)

	// assertions
	require.NoError(t, err)
}

func TestValidateVideo_UnhappyPath_VideoIsNil(t *testing.T) {
	// test
	err := ValidateVideo(nil)

	// assertions
	require.EqualError(t, err, ErrVideoIsNil.Error())
}

func TestValidateVideo_UnhappyPath_TitleIsInvalid(t *testing.T) {
	//fixture
	video := &model.Video{
		Title:       "",
		Description: "test description",
		Link:        "https://example.com/test",
		UserID:      1,
		Duration:    time.Duration(1),
		CreatedAt:   time.Now(),
	}

	// test
	err := ValidateVideo(video)

	// assertions
	require.EqualError(t, err, ErrTitleIsInvalid.Error())
}

func TestValidateVideo_UnhappyPath_DescriptionIsInvalid(t *testing.T) {
	// fixture
	video := &model.Video{
		Title:       "test title",
		Description: "",
		Link:        "https://example.com/test",
		UserID:      1,
		Duration:    time.Duration(1),
		CreatedAt:   time.Now(),
	}

	// test
	err := ValidateVideo(video)

	// assertions
	require.EqualError(t, err, ErrDescriptionIsInvalid.Error())
}

func TestValidateVideo_UnhappyPath_LinkIsInvalid(t *testing.T) {
	// fixture
	video := &model.Video{
		Title:       "test title",
		Description: "test description",
		Link:        "",
		UserID:      1,
		Duration:    time.Duration(1),
		CreatedAt:   time.Now(),
	}

	// test
	err := ValidateVideo(video)

	// assertions
	require.EqualError(t, err, ErrLinkIsInvalid.Error())
}

func TestValidateVideo_UnhappyPath_UserIDIsInvalid(t *testing.T) {
	// fixture
	video := &model.Video{
		Title:       "test title",
		Description: "test description",
		Link:        "https://example.com/test",
		UserID:      0,
		Duration:    time.Duration(1),
		CreatedAt:   time.Now(),
	}

	// test
	err := ValidateVideo(video)

	// assertions
	require.EqualError(t, err, ErrVideoUserIdIsInvalid.Error())
}

func TestValidateVideo_UnhappyPath_VideoCreatedAtIsInvalid(t *testing.T) {
	// fixture
	video := &model.Video{
		Title:       "test title",
		Description: "test description",
		Link:        "https://example.com/test",
		UserID:      1,
		Duration:    time.Duration(1),
		CreatedAt:   time.Time{},
	}

	// test
	err := ValidateVideo(video)

	// assertions
	require.EqualError(t, err, ErrVideoCreatedAtIsInvalid.Error())
}

func TestValidateVideo_UnhappyPath_VideoDurationsIsInvalid(t *testing.T) {
	// fixture
	video := &model.Video{
		Title:       "test title",
		Description: "test description",
		Link:        "https://example.com/test",
		UserID:      1,
		Duration:    time.Duration(0),
		CreatedAt:   time.Now(),
	}

	// test
	err := ValidateVideo(video)

	// assertions
	require.EqualError(t, err, ErrDurationIsInvalid.Error())
}

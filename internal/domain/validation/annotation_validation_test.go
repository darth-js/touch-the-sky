package validation

import (
	"testing"
	"time"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
	"github.com/stretchr/testify/require"
)

func TestValidateAnnotation_HappyPath(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "test type",
		UserID:    1,
		VideoID:   1,
		StartTime: time.Duration(1) * time.Minute,
		EndTime:   time.Duration(2) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.NoError(t, err)
}

func TestValidateAnnotation_UnhappyPath_AnnotationIsNil(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute

	// test
	err := ValidateAnnotation(nil, videoDuration)

	// assertions
	require.EqualError(t, err, ErrAnnotationIsNil.Error())
}

func TestValidateAnnotation_UnhappyPath_NoteIsInvalid(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "",
		Type:      "test type",
		UserID:    1,
		VideoID:   1,
		StartTime: time.Duration(1) * time.Minute,
		EndTime:   time.Duration(2) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrNoteIsInvalid.Error())
}

func TestValidateAnnotation_UnhappyPath_TypeIsInvalid(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "",
		UserID:    1,
		VideoID:   1,
		StartTime: time.Duration(1) * time.Minute,
		EndTime:   time.Duration(2) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrTypeIsInvalid.Error())
}

func TestValidateAnnotation_UnhappyPath_AnnotationUserIdIsInvalid(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "test type",
		UserID:    0,
		VideoID:   1,
		StartTime: time.Duration(1) * time.Minute,
		EndTime:   time.Duration(2) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrAnnotationUserIdIsInvalid.Error())
}

func TestValidateAnnotation_UnhappyPath_AnnotationVideoIdIdIsInvalid(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "test type",
		UserID:    1,
		VideoID:   0,
		StartTime: time.Duration(1) * time.Minute,
		EndTime:   time.Duration(2) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrAnnotationVideoIdIdIsInvalid.Error())
}

func TestValidateAnnotation_UnhappyPath_StartimeIsInvalid(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "test type",
		UserID:    1,
		VideoID:   1,
		StartTime: time.Duration(0),
		EndTime:   time.Duration(2) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrStartimeIsInvalid.Error())
}

func TestValidateAnnotation_UnhappyPath_EndtimeIsInvalid(t *testing.T) {
	//fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "test type",
		UserID:    1,
		VideoID:   1,
		StartTime: time.Duration(1) * time.Minute,
		EndTime:   time.Duration(0),
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrEndtimeIsInvalid.Error())
}

func TestValidateAnnotation_UnhappyPath_EndtimeIsLessThanStarttime(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "test type",
		UserID:    1,
		VideoID:   1,
		StartTime: time.Duration(2) * time.Minute,
		EndTime:   time.Duration(1) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrEndtimeIsInvalid.Error())
}

func TestValidateAnnotation_UnhappyPath_StarttimeIsGreaterThanVideoDuration(t *testing.T) {
	// fixture
	videoDuration := time.Duration(10) * time.Minute
	annotation := &model.Annotation{
		Note:      "test note",
		Type:      "test type",
		UserID:    1,
		VideoID:   1,
		StartTime: time.Duration(11) * time.Minute,
		EndTime:   time.Duration(12) * time.Minute,
	}

	// test
	err := ValidateAnnotation(annotation, videoDuration)

	// assertions
	require.EqualError(t, err, ErrStartimeIsInvalid.Error())
}

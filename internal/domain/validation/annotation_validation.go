package validation

import (
	"fmt"
	"time"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
)

var (
	ErrAnnotationIsNil              = fmt.Errorf("annotation is nil")
	ErrNoteIsInvalid                = fmt.Errorf("note is invalid")
	ErrTypeIsInvalid                = fmt.Errorf("type is invalid")
	ErrAnnotationUserIdIsInvalid    = fmt.Errorf("user id is invalid")
	ErrAnnotationVideoIdIdIsInvalid = fmt.Errorf("video id is invalid")
	ErrStartimeIsInvalid            = fmt.Errorf("startime is invalid")
	ErrEndtimeIsInvalid             = fmt.Errorf("endtime is invalid")

	AnnotationValidationErrors = map[error]bool{
		ErrAnnotationIsNil:              true,
		ErrNoteIsInvalid:                true,
		ErrTypeIsInvalid:                true,
		ErrAnnotationUserIdIsInvalid:    true,
		ErrAnnotationVideoIdIdIsInvalid: true,
		ErrStartimeIsInvalid:            true,
		ErrEndtimeIsInvalid:             true,
	}
)

func ValidateAnnotation(annotation *model.Annotation, videoDuration time.Duration) error {
	if annotation == nil {
		return ErrAnnotationIsNil
	}
	if annotation.Note == "" {
		return ErrNoteIsInvalid
	}

	if annotation.Type == "" {
		return ErrTypeIsInvalid
	}

	if annotation.UserID == 0 {
		return ErrAnnotationUserIdIsInvalid
	}

	if annotation.VideoID == 0 {
		return ErrAnnotationVideoIdIdIsInvalid
	}

	if annotation.StartTime == time.Duration(0) {
		return ErrStartimeIsInvalid
	}

	if annotation.EndTime == time.Duration(0) {
		return ErrEndtimeIsInvalid
	}

	if annotation.EndTime.Milliseconds() <= annotation.StartTime.Milliseconds() {
		return ErrEndtimeIsInvalid
	}

	if videoDuration.Milliseconds() < annotation.EndTime.Milliseconds() {
		return ErrStartimeIsInvalid
	}

	return nil
}

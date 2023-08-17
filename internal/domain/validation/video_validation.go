package validation

import (
	"fmt"
	"time"

	"github.com/juliocnsouzadev/go-videos-api/internal/domain/model"
)

var (
	ErrVideoIsNil              = fmt.Errorf("video is nil")
	ErrTitleIsInvalid          = fmt.Errorf("title is invalid")
	ErrDescriptionIsInvalid    = fmt.Errorf("description is invalid")
	ErrLinkIsInvalid           = fmt.Errorf("link is invalid")
	ErrVideoUserIdIsInvalid    = fmt.Errorf("user id is invalid")
	ErrDurationIsInvalid       = fmt.Errorf("duration is invalid")
	ErrVideoCreatedAtIsInvalid = fmt.Errorf("created_at is invalid")

	VideoValidationErrors = map[error]bool{
		ErrVideoIsNil:              true,
		ErrTitleIsInvalid:          true,
		ErrDescriptionIsInvalid:    true,
		ErrLinkIsInvalid:           true,
		ErrVideoUserIdIsInvalid:    true,
		ErrDurationIsInvalid:       true,
		ErrVideoCreatedAtIsInvalid: true,
	}
)

func ValidateVideo(video *model.Video) error {
	if video == nil {
		return ErrVideoIsNil
	}
	if video.Title == "" {
		return ErrTitleIsInvalid
	}

	if video.Description == "" {
		return ErrDescriptionIsInvalid
	}

	if video.Link == "" {
		return ErrLinkIsInvalid
	}

	if video.UserID == 0 {
		return ErrVideoUserIdIsInvalid
	}

	if video.Duration == time.Duration(0) {
		return ErrDurationIsInvalid
	}

	if video.CreatedAt.IsZero() {
		return ErrVideoCreatedAtIsInvalid
	}
	return nil
}

package model

import "time"

type Annotation struct {
	ID        int           `db:"id"`
	VideoID   int           `db:"video_id"`
	UserID    int           `db:"user_id"`
	StartTime time.Duration `db:"start_time"`
	EndTime   time.Duration `db:"end_time"`
	Type      string        `db:"type"`
	Note      string        `db:"note"`
}

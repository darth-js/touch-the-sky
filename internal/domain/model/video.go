package model

import "time"

type Video struct {
	ID          int           `db:"id"`
	UserID      int           `db:"user_id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
	Link        string        `db:"link"`
	Duration    time.Duration `db:"duration"`
	CreatedAt   time.Time     `db:"created_at"`
}

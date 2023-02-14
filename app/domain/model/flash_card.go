package model

import "time"

type FlashCard struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Front     string    `db:"front"`
	Back      string    `db:"back"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

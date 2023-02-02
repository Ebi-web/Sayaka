package model

import "time"

type User struct {
	ID          int64     `db:"id"`
	LINEUserID  string    `db:"line_user_id"`
	DisplayName string    `db:"display_name"`
	PhotoURL    string    `db:"photo_url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

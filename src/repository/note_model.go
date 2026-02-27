package repository

import "time"

type noteModel struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	UpdatedAt time.Time `db:"updated_at"`
}

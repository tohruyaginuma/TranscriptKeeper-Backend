package model

import "time"

type NoteModel struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	UpdatedAt time.Time `db:"updated_at"`
}

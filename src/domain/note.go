package domain

import (
	"errors"
	"strconv"
	"time"
)

const (
	MinNoteID          = 1
	MinNoteTitleLength = 1
)

type NoteID int64

var (
	ErrInvalidNoteID    = errors.New("note_id must be greater than 0")
	ErrInvalidNoteTitle = errors.New("note_title must be greater than 0")
	ErrNoteNotFound     = errors.New("note not found")
)

func NewNoteID(id int64) (NoteID, error) {
	if id < MinNoteID {
		return NoteID(0), ErrInvalidNoteID
	}
	return NoteID(id), nil
}

func ParseNoteID(s string) (NoteID, error) {
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return NoteID(0), ErrInvalidNoteID
	}
	return NewNoteID(id)
}

type NoteTitle string

func NewNoteTitle(title string) (NoteTitle, error) {
	if len(title) < MinNoteTitleLength {
		return NoteTitle(""), ErrInvalidNoteTitle
	}
	return NoteTitle(title), nil
}

type Note struct {
	ID        NoteID
	Title     NoteTitle
	UserID    UserID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewNote(id NoteID, title NoteTitle, userID UserID) *Note {
	return &Note{
		ID:     id,
		Title:  title,
		UserID: userID,
	}
}

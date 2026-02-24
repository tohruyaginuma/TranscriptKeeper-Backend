package domain

import (
	"errors"
)

type NoteID int64

var (
	ErrInvalidNoteID = errors.New("note_id must be greater than 0")
	ErrInvalidTitle  = errors.New("title must be greater than 0")
)

func NewNoteID(id int64) (NoteID, error) {
	if id <= 0 {
		return NoteID(0), ErrInvalidNoteID
	}
	return NoteID(id), nil
}

type Title string

func NewTitle(title string) (Title, error) {
	if len(title) == 0 {
		return Title(""), ErrInvalidTitle
	}
	return Title(title), nil
}

type Note struct {
	ID     NoteID
	Title  Title
	UserID UserID
}

func NewNote(id NoteID, title Title, userID UserID) *Note {
	return &Note{
		ID:     id,
		Title:  title,
		UserID: userID,
	}
}

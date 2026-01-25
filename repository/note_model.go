package repository

import (
	"fmt"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type NoteModel struct {
	ID     int64  `db:"id"`
	Name   string `db:"name"`
	UserID int64  `db:"user_id"`
}

func (m *NoteModel) toDomain() (note *domain.Note, err error) {
	noteID, err := domain.NewNoteID(m.ID)
	if err != nil {
		return note, fmt.Errorf("failed to convert note id: %w", err)
	}

	noteName, err := domain.NewNoteName(m.Name)
	if err != nil {
		return note, fmt.Errorf("failed to convert note name: %w", err)
	}

	userID, err := domain.NewUserID(m.UserID)
	if err != nil {
		return note, fmt.Errorf("failed to convert user id: %w", err)
	}

	note = domain.NewNote(noteID, noteName, userID)

	return note, nil
}

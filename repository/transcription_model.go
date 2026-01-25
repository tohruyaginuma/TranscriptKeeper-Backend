package repository

import (
	"fmt"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type transcriptionModel struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
	NoteID  int64  `db:"note_id"`
}

func (m *transcriptionModel) toDomain() (transcription domain.Transcription, err error) {
	id, err := domain.NewTranscriptionID(m.ID)
	if err != nil {
		return domain.Transcription{}, fmt.Errorf("transcriptionID convert failed: %w", err)
	}

	content, err := domain.NewTranscriptionContent(m.Content)
	if err != nil {
		return domain.Transcription{}, fmt.Errorf("transcriptionContent convert failed: %w", err)
	}

	noteID, err := domain.NewNoteID(m.NoteID)
	if err != nil {
		return domain.Transcription{}, fmt.Errorf("transcription convert failed: %w", err)
	}

	transcriptionMemory := domain.NewTranscription(id, content, noteID)

	return *transcriptionMemory, nil
}

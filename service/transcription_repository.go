package service

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type TranscriptionRepository interface {
	Create(ctx context.Context, noteID domain.NoteID, content domain.TranscriptionContent) (transcriptionID domain.TranscriptionID, err error)
	ListByNoteID(ctx context.Context, noteID domain.NoteID, limit, offset int) (transcriptions []domain.Transcription, err error)
}

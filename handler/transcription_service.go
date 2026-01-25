package handler

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type TranscriptionService interface {
	Create(ctx context.Context, noteID domain.NoteID, content string) (transcriptionID domain.TranscriptionID, err error)
	ListByNoteID(ctx context.Context, noteID domain.NoteID, limit, offset int) (transcriptions []domain.Transcription, err error)
}

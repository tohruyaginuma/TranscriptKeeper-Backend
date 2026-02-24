package application

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type TranscriptRepository interface {
	Create(ctx context.Context, transcript *domain.Transcript) (transcript *domain.Transcript, err error)
	ListByNoteID(ctx context.Context, noteID domain.NoteID) (transcripts []*domain.Transcript, err error)
}

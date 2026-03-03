package handler

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type transcriptUsecase interface {
	Create(ctx context.Context, content domain.Content, noteID domain.NoteID) (transcriptID domain.TranscriptID, err error)
	ListByNoteID(ctx context.Context, noteID domain.NoteID) (transcripts []*domain.Transcript, err error)
}

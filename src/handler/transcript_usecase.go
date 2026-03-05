package handler

import (
	"context"
	"io"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/transcript"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type transcriptUsecase interface {
	Create(ctx context.Context, audio io.Reader, language string, noteID domain.NoteID) (transcriptResult transcript.TranscriptResult, err error)
	RetrieveByNoteID(ctx context.Context, noteID domain.NoteID) (transcript domain.Transcript, err error)
}

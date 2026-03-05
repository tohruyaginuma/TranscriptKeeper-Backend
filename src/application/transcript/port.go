package transcript

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type TranscriptRepository interface {
	Create(ctx context.Context, text int64, noteID domain.NoteID) (transcriptID domain.TranscriptID, err error)
	RetrieveByNoteID(ctx context.Context, noteID domain.NoteID) (transcript domain.Transcript, err error)
}

type TranscriptResult struct {
	TranscriptID domain.TranscriptID
}

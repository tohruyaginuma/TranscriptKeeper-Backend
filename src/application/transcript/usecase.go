package transcript

import (
	"context"
	"io"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/db/postgres"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/external"
)

type TranscriptUsecase struct {
	CloudflareClient     *external.CloudflareWorkersAIClient
	TranscriptRepository *postgres.TranscriptRepository
}

func NewTranscriptUsecase(cloudflareClient *external.CloudflareWorkersAIClient, transcriptRepository *postgres.TranscriptRepository) *TranscriptUsecase {
	return &TranscriptUsecase{
		CloudflareClient:     cloudflareClient,
		TranscriptRepository: transcriptRepository,
	}
}

func (s *TranscriptUsecase) Create(ctx context.Context, audio io.Reader, language string, noteID domain.NoteID) (transcriptResult TranscriptResult, err error) {
	result, err := s.CloudflareClient.Transcribe(ctx, audio, language)
	if err != nil {
		return TranscriptResult{}, err
	}

	transcriptID, err := s.TranscriptRepository.Create(ctx, result.Text, noteID)
	if err != nil {
		return TranscriptResult{}, err
	}

	return TranscriptResult{TranscriptID: transcriptID}, nil
}

func (s *TranscriptUsecase) RetrieveByNoteID(ctx context.Context, noteID domain.NoteID) (transcript domain.Transcript, err error) {
	transcript, err = s.TranscriptRepository.RetrieveByNoteID(ctx, noteID)
	if err != nil {
		return domain.Transcript{}, err
	}

	return transcript, nil
}

package service

import (
	"context"
	"fmt"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type transcriptionService struct {
	repo     TranscriptionRepository
	repoNote NoteRepository
}

func NewTranscriptionService(repo TranscriptionRepository, repoNote NoteRepository) *transcriptionService {
	return &transcriptionService{
		repo:     repo,
		repoNote: repoNote,
	}
}

func (s *transcriptionService) Create(ctx context.Context, noteID domain.NoteID, content string) (transcriptionID domain.TranscriptionID, err error) {
	exists, err := s.repoNote.Exists(ctx, noteID)
	if err != nil {
		return domain.TranscriptionID(0), fmt.Errorf("exists note failed: %w", err)
	}
	if !exists {
		return domain.TranscriptionID(0), ErrNoteNotFound
	}

	transcriptionContent, err := domain.NewTranscriptionContent(content)
	if err != nil {
		return domain.TranscriptionID(0), fmt.Errorf("convert content failed: %w", err)
	}

	transcriptionID, err = s.repo.Create(ctx, noteID, transcriptionContent)
	if err != nil {
		return domain.TranscriptionID(0), fmt.Errorf("create transcription failed: %w", err)
	}

	return transcriptionID, nil
}

func (s *transcriptionService) ListByNoteID(ctx context.Context, noteID domain.NoteID, limit, offset int) (transcriptions []domain.Transcription, err error) {
	exists, err := s.repoNote.Exists(ctx, noteID)
	if err != nil {
		return nil, fmt.Errorf("exists note failed: %w", err)
	}
	if !exists {
		return nil, ErrNoteNotFound
	}

	transcriptions, err = s.repo.ListByNoteID(ctx, noteID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list transcriptions failed: %w", err)
	}

	return transcriptions, nil
}

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/repository"
)

var (
	ErrNoteNotFound = errors.New("note not found")
)

type noteService struct {
	repo NoteRepository
}

func NewNoteService(repo NoteRepository) *noteService {
	return &noteService{
		repo: repo,
	}
}

func (s *noteService) Create(ctx context.Context, userID domain.UserID, name string) (domain.NoteID, error) {
	noteName, err := domain.NewNoteName(name)
	if err != nil {
		return domain.NoteID(0), fmt.Errorf("create note: new note name: %w", err)
	}
	noteID, err := s.repo.Create(ctx, userID, noteName)
	if err != nil {
		return domain.NoteID(0), fmt.Errorf("create note: %w", err)
	}

	return noteID, nil
}

func (s *noteService) List(ctx context.Context, userID domain.UserID, limit, offset int) (notes []domain.Note, count int, err error) {
	notes, err = s.repo.List(ctx, userID, limit, offset)
	if err != nil {
		return notes, 0, fmt.Errorf("failed to list notes: %w", err)
	}

	count, err = s.repo.Count(ctx, userID)
	if err != nil {
		return notes, 0, fmt.Errorf("failed to count notes: %w", err)
	}

	return notes, count, nil
}

func (s *noteService) Delete(ctx context.Context, noteID domain.NoteID) (err error) {
	if err = s.repo.Delete(ctx, noteID); err != nil {
		if errors.Is(err, repository.ErrNoteNotFound) {
			return fmt.Errorf("note not found: %w", ErrNoteNotFound)
		}

		return fmt.Errorf("delete note failed: %w", err)
	}

	return nil
}

func (s *noteService) Update(ctx context.Context, noteID domain.NoteID, name string, userID domain.UserID) (err error) {
	noteName, err := domain.NewNoteName(name)
	if err != nil {
		return fmt.Errorf("convert domain noteName failed: %w", err)
	}
	note := domain.NewNote(noteID, noteName, userID)
	_, err = s.repo.Update(ctx, note)
	if err != nil {
		if errors.Is(err, repository.ErrNoteNotFound) {
			return ErrNoteNotFound
		}
		return fmt.Errorf("failed to update: %w", err)
	}

	return nil
}

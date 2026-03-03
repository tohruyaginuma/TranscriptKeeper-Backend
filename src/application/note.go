package application

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type NoteUsecase struct {
	noteRepository NoteRepository
}

func NewNoteUsecase(noteRepository NoteRepository) *NoteUsecase {
	return &NoteUsecase{noteRepository: noteRepository}
}

func (s *NoteUsecase) Create(ctx context.Context, title domain.NoteTitle, userID domain.UserID) (noteID domain.NoteID, err error) {
	return s.noteRepository.Create(ctx, title, userID)
}

func (s *NoteUsecase) Update(ctx context.Context, noteID domain.NoteID, title domain.NoteTitle) (err error) {
	return s.noteRepository.Update(ctx, noteID, title)
}

func (s *NoteUsecase) Delete(ctx context.Context, noteID domain.NoteID) (err error) {
	return s.noteRepository.Delete(ctx, noteID)
}

func (s *NoteUsecase) ListByUserID(ctx context.Context, userID domain.UserID) (notes []*NoteListItem, err error) {
	return s.noteRepository.ListByUserID(ctx, userID)
}

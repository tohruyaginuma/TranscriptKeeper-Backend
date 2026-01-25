package service

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type NoteRepository interface {
	Create(ctx context.Context, userID domain.UserID, name domain.NoteName) (domain.NoteID, error)
	List(ctx context.Context, userID domain.UserID, limit, offset int) (notes []domain.Note, err error)
	Delete(ctx context.Context, noteID domain.NoteID) (err error)
	Update(ctx context.Context, note *domain.Note) (noteID domain.NoteID, err error)
	Count(ctx context.Context, userID domain.UserID) (count int, err error)
	Exists(ctx context.Context, noteID domain.NoteID) (exists bool, err error)
}

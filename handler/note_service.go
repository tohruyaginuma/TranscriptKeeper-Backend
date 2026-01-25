package handler

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type NoteService interface {
	Create(ctx context.Context, userID domain.UserID, name string) (domain.NoteID, error)
	List(ctx context.Context, userID domain.UserID, limit, offset int) (notes []domain.Note, count int, err error)
	Update(ctx context.Context, noteID domain.NoteID, name string, userID domain.UserID) (err error)
	Delete(ctx context.Context, noteID domain.NoteID) (err error)
}

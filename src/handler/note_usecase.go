package handler

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type noteUsecase interface {
	Create(ctx context.Context, title domain.Title, userID domain.UserID) (note *domain.Note, err error)
	// Update(ctx context.Context, note *domain.Note) (note *domain.Note, err error)
	// Delete(ctx context.Context, id domain.NoteID) (err error)
	// ListByUserID(ctx context.Context, userID domain.UserID) (notes []*domain.Note, err error)
}

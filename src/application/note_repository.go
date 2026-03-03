package application

import (
	"context"
	"time"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type NoteRepository interface {
	Create(ctx context.Context, title domain.NoteTitle, userID domain.UserID) (noteID domain.NoteID, err error)
	Update(ctx context.Context, noteID domain.NoteID, title domain.NoteTitle) (err error)
	Delete(ctx context.Context, noteID domain.NoteID) (err error)
	ListByUserID(ctx context.Context, userID domain.UserID) (notes []*NoteListItem, err error)
}

type NoteListItem struct {
	ID        int64
	Title     string
	UpdatedAt time.Time
}

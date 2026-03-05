package note

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/note"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type noteUsecase interface {
	Create(ctx context.Context, title domain.NoteTitle, userID domain.UserID) (noteID domain.NoteID, err error)
	Update(ctx context.Context, noteID domain.NoteID, title domain.NoteTitle) (err error)
	Delete(ctx context.Context, noteID domain.NoteID) (err error)
	ListByUserID(ctx context.Context, userID domain.UserID) (notes []*note.NoteListItem, err error)
}

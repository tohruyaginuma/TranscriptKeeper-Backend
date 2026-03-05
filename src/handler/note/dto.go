package note

import (
	"time"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/note"
)

type noteCreateRequest struct {
	Title string `json:"title" validate:"required"`
}

type noteUpdateRequest struct {
	Title string `json:"title" validate:"required"`
}
type noteResponseItem struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
}

type noteResponse struct {
	Notes []noteResponseItem `json:"notes"`
}

func newNoteResponse(notes []*note.NoteListItem) *noteResponse {
	response := &noteResponse{
		Notes: make([]noteResponseItem, len(notes)),
	}
	for i, note := range notes {
		response.Notes[i] = noteResponseItem{
			ID:        note.ID,
			Title:     note.Title,
			UpdatedAt: note.UpdatedAt,
		}
	}
	return response
}

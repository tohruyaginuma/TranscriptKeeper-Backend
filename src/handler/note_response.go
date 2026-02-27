package handler

import (
	"time"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type noteResponseItem struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
}

type noteResponse struct {
	Notes []noteResponseItem `json:"notes"`
}

func newNoteResponse(notes []*domain.Note) *noteResponse {
	response := &noteResponse{
		Notes: make([]noteResponseItem, len(notes)),
	}
	for i, note := range notes {
		response.Notes[i] = noteResponseItem{
			ID:        int64(note.ID),
			Title:     string(note.Title),
			UpdatedAt: note.UpdatedAt,
		}
	}
	return response
}

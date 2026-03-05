package domain_test

import (
	"errors"
	"testing"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

func TestNewNoteID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   int64
		want    domain.NoteID
		wantErr error
	}{
		{
			name:    "valid id",
			input:   1,
			want:    domain.NoteID(1),
			wantErr: nil,
		},
		{
			name:    "zero id",
			input:   0,
			want:    domain.NoteID(0),
			wantErr: domain.ErrInvalidNoteID,
		},
		{
			name:    "negative id",
			input:   -1,
			want:    domain.NoteID(0),
			wantErr: domain.ErrInvalidNoteID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewNoteID(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewNoteID() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("NewNoteID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNoteTitle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    domain.NoteTitle
		wantErr error
	}{
		{
			name:    "valid title",
			input:   "Meeting memo",
			want:    domain.NoteTitle("Meeting memo"),
			wantErr: nil,
		},
		{
			name:    "empty title",
			input:   "",
			want:    domain.NoteTitle(""),
			wantErr: domain.ErrInvalidNoteTitle,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewNoteTitle(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewNoteTitle() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("NewNoteTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNote(t *testing.T) {
	t.Parallel()

	id := domain.NoteID(1)
	title := domain.NoteTitle("Meeting memo")
	userID := domain.UserID(10)

	got := domain.NewNote(id, title, userID)
	if got == nil {
		t.Fatal("NewNote() returned nil")
	}

	if got.ID != id {
		t.Fatalf("NewNote().ID = %v, want %v", got.ID, id)
	}

	if got.Title != title {
		t.Fatalf("NewNote().Title = %v, want %v", got.Title, title)
	}

	if got.UserID != userID {
		t.Fatalf("NewNote().UserID = %v, want %v", got.UserID, userID)
	}

	if !got.CreatedAt.IsZero() {
		t.Fatalf("NewNote().CreatedAt = %v, want zero value", got.CreatedAt)
	}

	if !got.UpdatedAt.IsZero() {
		t.Fatalf("NewNote().UpdatedAt = %v, want zero value", got.UpdatedAt)
	}
}

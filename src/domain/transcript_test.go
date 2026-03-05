package domain_test

import (
	"errors"
	"testing"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

func TestNewTranscriptID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   int64
		want    domain.TranscriptID
		wantErr error
	}{
		{
			name:    "valid id",
			input:   1,
			want:    domain.TranscriptID(1),
			wantErr: nil,
		},
		{
			name:    "zero id",
			input:   0,
			want:    domain.TranscriptID(0),
			wantErr: domain.ErrInvalidTranscriptID,
		},
		{
			name:    "negative id",
			input:   -1,
			want:    domain.TranscriptID(0),
			wantErr: domain.ErrInvalidTranscriptID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewTranscriptID(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewTranscriptID() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("NewTranscriptID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    domain.Content
		wantErr error
	}{
		{
			name:    "valid content",
			input:   "hello world",
			want:    domain.Content("hello world"),
			wantErr: nil,
		},
		{
			name:    "empty content",
			input:   "",
			want:    domain.Content(""),
			wantErr: domain.ErrInvalidContent,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewContent(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewContent() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("NewContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTranscript(t *testing.T) {
	t.Parallel()

	id := domain.TranscriptID(1)
	text := "transcribed text"
	noteID := domain.NoteID(3)

	got := domain.NewTranscript(id, text, noteID)
	if got == nil {
		t.Fatal("NewTranscript() returned nil")
	}

	if got.ID != id {
		t.Fatalf("NewTranscript().ID = %v, want %v", got.ID, id)
	}

	if got.Text != text {
		t.Fatalf("NewTranscript().Text = %v, want %v", got.Text, text)
	}

	if got.NoteID != noteID {
		t.Fatalf("NewTranscript().NoteID = %v, want %v", got.NoteID, noteID)
	}
}

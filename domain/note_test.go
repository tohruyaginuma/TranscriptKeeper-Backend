package domain_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

func TestNewNoteID_OK(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int64
		want domain.NoteID
	}{
		{"id=1", 1, domain.NoteID(1)},
		{"id=50", 50, domain.NoteID(50)},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := domain.NewNoteID(tt.in)
			if err != nil {
				t.Fatalf("NewNoteID(%d) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("NewNoteID(%d) got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}

}
func TestNewNoteID_NG_Invalid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int64
	}{
		{"id=0", 0},
		{"id=-100", -100},
		{"id=-999", -999},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := domain.NewNoteID(tt.in)
			if err == nil {
				t.Fatalf("NoteID(%d) expected error. got nil", tt.in)
			}
			if !errors.Is(err, domain.ErrNoteInvalidArgument) {
				t.Fatalf("NoteID(%d) error = %v, want ErrNoteInvalidArgument", tt.in, err)
			}
		})
	}
}
func TestNewNoteName_OK(t *testing.T) {
	t.Parallel()

	maxText := strings.Repeat("a", 255)

	tests := []struct {
		name string
		in   string
		want domain.NoteName
	}{
		{"name=meeting", "meeting", domain.NoteName("meeting")},
		{"name=maximum", maxText, domain.NoteName(maxText)},
		{"name=trim", "  meeting  ", domain.NoteName("meeting")},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res, err := domain.NewNoteName(tt.in)
			if err != nil {
				t.Fatalf("NewNoteName(%v) returned error: %v", tt.in, err)
			}
			if res != tt.want {
				t.Fatalf("NewNoteName(%v) = %v, want %v", tt.in, res, tt.want)
			}
		})
	}
}

func TestNewNoteName_NG_Invalid(t *testing.T) {
	t.Parallel()

	overText := strings.Repeat("a", 256)

	tests := []struct {
		name string
		in   string
	}{
		{"name empty", ""},
		{"name empty with space", "  "},
		{"name over length", overText},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := domain.NewNoteName(tt.in)
			if err == nil {
				t.Fatalf("NoteName(%v) expected err, got nil", tt.in)
			}
			if !errors.Is(err, domain.ErrNoteInvalidArgument) {
				t.Fatalf("unexpected error, expected ErrNoteInvalidArgument")
			}
		})
	}
}

func TestNewNote_OK(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		id       int64
		noteName string
		userID   int64
	}{
		{"basic", 1, "meeting", 1},
		{"different values", 100, "project", 50},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			noteID, _ := domain.NewNoteID(tt.id)
			noteName, _ := domain.NewNoteName(tt.noteName)
			userID, _ := domain.NewUserID(tt.userID)

			note := domain.NewNote(noteID, noteName, userID)

			if note.ID() != noteID {
				t.Fatalf("note.ID() = %v, want: %v", note.ID(), noteID)
			}
			if note.Name() != noteName {
				t.Fatalf("note.Name() = %v, want: %v", note.Name(), noteName)
			}
			if note.UserID() != userID {
				t.Fatalf("note.UserID() = %v, want: %v", note.UserID(), userID)
			}
		})
	}
}

package domain_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

func TestNewUserID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   int64
		want    domain.UserID
		wantErr error
	}{
		{
			name:    "valid id",
			input:   1,
			want:    domain.UserID(1),
			wantErr: nil,
		},
		{
			name:    "zero id",
			input:   0,
			want:    domain.UserID(0),
			wantErr: domain.ErrInvalidUserID,
		},
		{
			name:    "negative id",
			input:   -1,
			want:    domain.UserID(0),
			wantErr: domain.ErrInvalidUserID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewUserID(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewUserID() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("NewUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFirebaseUID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    domain.FirebaseUID
		wantErr error
	}{
		{
			name:    "valid uid",
			input:   "firebase-uid-1",
			want:    domain.FirebaseUID("firebase-uid-1"),
			wantErr: nil,
		},
		{
			name:    "max length uid",
			input:   strings.Repeat("a", 255),
			want:    domain.FirebaseUID(strings.Repeat("a", 255)),
			wantErr: nil,
		},
		{
			name:    "empty uid",
			input:   "",
			want:    domain.FirebaseUID(""),
			wantErr: domain.ErrInvalidFirebaseUID,
		},
		{
			name:    "too long uid",
			input:   strings.Repeat("a", 256),
			want:    domain.FirebaseUID(""),
			wantErr: domain.ErrInvalidFirebaseUID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := domain.NewFirebaseUID(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewFirebaseUID() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("NewFirebaseUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	t.Parallel()

	id := domain.UserID(1)
	firebaseUID := domain.FirebaseUID("firebase-uid-1")

	got := domain.NewUser(id, firebaseUID)
	if got == nil {
		t.Fatal("NewUser() returned nil")
	}

	if got.ID != id {
		t.Fatalf("NewUser().ID = %v, want %v", got.ID, id)
	}

	if got.FirebaseUID != firebaseUID {
		t.Fatalf("NewUser().FirebaseUID = %v, want %v", got.FirebaseUID, firebaseUID)
	}

	if !got.CreatedAt.IsZero() {
		t.Fatalf("NewUser().CreatedAt = %v, want zero value", got.CreatedAt)
	}

	if !got.UpdatedAt.IsZero() {
		t.Fatalf("NewUser().UpdatedAt = %v, want zero value", got.UpdatedAt)
	}
}

package domain

import (
	"errors"
	"testing"
)

func TeNewUserID_OK(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int64
		want UserID
	}{
		{"id=1", 1, UserID(1)},
		{"id=42", 42, UserID(42)},
	}

	for _, tt := range tests {
		// For Pararel
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewUserID(tt.in)
			if err != nil {
				t.Fatalf("NewUserID(%d) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("NewUserID(%d) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestNewUserID_Invalid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int64
	}{
		{"id=0", 0},
		{"id=-1", -1},
		{"id=-999", -999},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewUserID(tt.in)
			if err == nil {
				t.Fatalf("NewUserID(%d) expected error, got nil", tt.in)
			}
			if !errors.Is(err, ErrInvalidUserID) {
				t.Fatalf("NewUserID(%d) error = %v, want ErrInvalidUserID", tt.in, err)
			}
		})
	}

}

func TestNewUserName_OK(t *testing.T) {
	t.Parallel()

	got, err := NewUserName("Tohru")
	if err != nil {
		t.Fatalf("NewUserName returned error: %v", err)
	}
	if got.String() != "Tohru" {
		t.Fatalf("userName.String() = %q, want :%q", got.String(), "Tohru")
	}
}

func TestNewUserName_TrimsSpace(t *testing.T) {
	t.Parallel()

	got, err := NewUserName("  Tohru  ")
	if err != nil {
		t.Fatalf("NewUserName returned error: %v", err)
	}
	if got.String() != "Tohru" {
		t.Fatalf("UserName.String() = %q, want %q", got.String(), "Tohru")
	}
}

func TestNewUserName_Empty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
	}{
		{"empty", ""},
		{"space", "  "},
		{"tabs/newlines", "\n\t"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewUserName(tt.in)
			if err == nil {
				t.Fatalf("NewUserName(%q) expected err, got nil", tt.in)
			}
			if !errors.Is(err, ErrNameEmpty) {
				t.Fatalf("NewUserName(%q) error = %v, want ErrNameEmpty", tt.in, err)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	t.Parallel()

	id, err := NewUserID(10)
	if err != nil {
		t.Fatalf("NewuserID returned error: %v", err)
	}

	name, err := NewUserName("  Alice  ")
	if err != nil {
		t.Fatalf("NewUserName returned error: %v", err)
	}

	user := NewUser(id, name)

	if user.ID() != id {
		t.Fatalf("user.ID() = %v, want %v", user.ID(), id)
	}
	if user.Name().String() != "Alice" {
		t.Fatalf("User.Name().String() = %q, want %q", user.Name().String(), "Alice")
	}
}
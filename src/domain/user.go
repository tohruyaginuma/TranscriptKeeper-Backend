package domain

import (
	"errors"
)

var (
	ErrInvalidGoogleID = errors.New("google_id must be 255 or less characters long")
	ErrInvalidUserID   = errors.New("user_id must be greater than 0")
)

type UserID int64

func NewUserID(id int64) (UserID, error) {
	if id <= 0 {
		return UserID(0), ErrInvalidUserID
	}

	return UserID(id), nil
}

type GoogleID string

func NewGoogleID(id string) (GoogleID, error) {
	if len(id) > 255 {
		return GoogleID(""), ErrInvalidGoogleID
	}

	return GoogleID(id), nil
}

type User struct {
	ID       UserID
	GoogleID GoogleID
}

func NewUser(id UserID, googleID GoogleID) *User {
	return &User{
		ID:       id,
		GoogleID: googleID,
	}
}

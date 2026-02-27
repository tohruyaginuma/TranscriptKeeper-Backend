package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidFirebaseUID   = errors.New("firebase_uid must be 255 or less characters long")
	ErrInvalidUserID        = errors.New("user_id must be greater than 0")
	ErrDuplicateFirebaseUID = errors.New("duplicate firebase_uid")
	ErrUserNotFound         = errors.New("user not found")
)

type UserID int64

func NewUserID(id int64) (UserID, error) {
	if id <= 0 {
		return UserID(0), ErrInvalidUserID
	}

	return UserID(id), nil
}

type FirebaseUID string

func NewFirebaseUID(uid string) (FirebaseUID, error) {
	if len(uid) == 0 || len(uid) > 255 {
		return FirebaseUID(""), ErrInvalidFirebaseUID
	}

	return FirebaseUID(uid), nil
}

type User struct {
	ID          UserID
	FirebaseUID FirebaseUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser(id UserID, firebaseUID FirebaseUID) *User {
	return &User{
		ID:          id,
		FirebaseUID: firebaseUID,
	}
}

package domain

import (
	"errors"
	"strings"
)

var (
	ErrNameEmpty     = errors.New("name is empty")
	ErrInvalidUserID = errors.New("invalid userID")
)

type UserID int64

func (u *UserID) Value() int64 { return int64(*u) }

func NewUserID(id int64) (UserID, error) {
	if id <= 0 {
		return 0, ErrInvalidUserID
	}

	return UserID(id), nil
}

type UserName string

func (n UserName) String() string { return string(n) }

func NewUserName(name string) (UserName, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return "", ErrNameEmpty
	}

	return UserName(name), nil
}

type User struct {
	id   UserID
	name UserName
}

func (u User) ID() UserID     { return u.id }
func (u User) Name() UserName { return u.name }

func NewUser(id UserID, name UserName) User {
	// you can validate here.

	return User{
		id:   id,
		name: name,
	}
}

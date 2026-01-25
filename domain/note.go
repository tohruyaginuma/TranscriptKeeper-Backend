package domain

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

var (
	ErrNoteInvalidArgument = errors.New("Invalid Argument")
)

type NoteID int64

func (n NoteID) Value() int64 { return int64(n) }

func NewNoteID(id int64) (NoteID, error) {
	if id <= 0 {
		return NoteID(0), fmt.Errorf("%w: id must be positive: %d", ErrNoteInvalidArgument, id)
	}

	return NoteID(id), nil
}

type NoteName string

func (v NoteName) String() string { return string(v) }

func NewNoteName(name string) (NoteName, error) {
	const MaxNoteNameLength = 255

	name = strings.TrimSpace(name)

	if name == "" {
		return NoteName(""), fmt.Errorf("%w: name is not allowed empty: %v", ErrNoteInvalidArgument, name)
	}
	if utf8.RuneCountInString(name) > MaxNoteNameLength {
		return NoteName(""), fmt.Errorf("%w: name exceeds maximum length of %v ", ErrNoteInvalidArgument, MaxNoteNameLength)
	}

	return NoteName(name), nil
}

type Note struct {
	id     NoteID
	name   NoteName
	userID UserID
}

func (n *Note) Name() NoteName { return n.name }
func (n *Note) ID() NoteID     { return n.id }
func (n *Note) UserID() UserID { return n.userID }

func NewNote(id NoteID, name NoteName, userID UserID) *Note {
	// Validation if needed
	return &Note{
		id:     id,
		name:   name,
		userID: userID,
	}
}

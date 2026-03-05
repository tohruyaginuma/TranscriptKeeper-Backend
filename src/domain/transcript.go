package domain

import (
	"errors"
)

const (
	MinTranscriptID            = 1
	MinTranscriptContentLength = 1
)

var (
	ErrInvalidTranscriptID = errors.New("transcript_id must be greater than 0")
	ErrInvalidContent      = errors.New("content must be greater than 0")
)

type TranscriptID int64

func NewTranscriptID(id int64) (TranscriptID, error) {
	if id < MinTranscriptID {
		return TranscriptID(0), ErrInvalidTranscriptID
	}
	return TranscriptID(id), nil
}

type Content string

func NewContent(content string) (Content, error) {
	if len(content) < MinTranscriptContentLength {
		return Content(""), ErrInvalidContent
	}
	return Content(content), nil
}

type Transcript struct {
	ID     TranscriptID
	Text   string
	NoteID NoteID
}

func NewTranscript(id TranscriptID, text string, noteID NoteID) *Transcript {
	return &Transcript{
		ID:     id,
		Text:   text,
		NoteID: noteID,
	}
}

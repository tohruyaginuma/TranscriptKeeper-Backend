package domain

import (
	"errors"
)

var (
	ErrInvalidTranscriptID = errors.New("transcript_id must be greater than 0")
	ErrInvalidContent      = errors.New("content must be greater than 0")
)

type TranscriptID int64

func NewTranscriptID(id int64) (TranscriptID, error) {
	if id <= 0 {
		return TranscriptID(0), ErrInvalidTranscriptID
	}
	return TranscriptID(id), nil
}

type Content string

func NewContent(content string) (Content, error) {
	if len(content) == 0 {
		return Content(""), ErrInvalidContent
	}
	return Content(content), nil
}

type Transcript struct {
	ID      TranscriptID
	Content Content
	NoteID  NoteID
}

func NewTranscript(id TranscriptID, content Content, noteID NoteID) *Transcript {
	return &Transcript{
		ID:      id,
		Content: content,
		NoteID:  noteID,
	}
}

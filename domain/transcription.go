package domain

import (
	"errors"
	"strings"
	"unicode/utf8"
)

const maxContent = 1000

var (
	ErrTranscriptionInvalidArgument = errors.New("Transcription Invalid Argument")
)

type TranscriptionID int64

func (t TranscriptionID) Value() int64 { return int64(t) }

func NewTranscriptionID(n int64) (TranscriptionID, error) {
	if n <= 0 {
		return TranscriptionID(0), ErrTranscriptionInvalidArgument
	}

	return TranscriptionID(n), nil
}

type TranscriptionContent string

func (t TranscriptionContent) String() string { return string(t) }

func NewTranscriptionContent(c string) (TranscriptionContent, error) {
	c = strings.TrimSpace(c)

	if c == "" {
		return TranscriptionContent(""), ErrTranscriptionInvalidArgument
	}
	if utf8.RuneCountInString(c) > maxContent {
		return TranscriptionContent(""), ErrTranscriptionInvalidArgument
	}

	return TranscriptionContent(c), nil
}

type Transcription struct {
	id      TranscriptionID
	content TranscriptionContent
	noteID  NoteID
}

func (t *Transcription) ID() TranscriptionID           { return t.id }
func (t *Transcription) Content() TranscriptionContent { return t.content }
func (t *Transcription) NoteID() NoteID                { return t.noteID }

func NewTranscription(id TranscriptionID, content TranscriptionContent, noteID NoteID) *Transcription {
	return &Transcription{
		id:      id,
		content: content,
		noteID:  noteID,
	}
}

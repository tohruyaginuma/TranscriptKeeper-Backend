package domain

type TranscriptRepository interface {
	Create(transcript *Transcript) (*Transcript, error)
	ListByNoteID(noteID NoteID) ([]*Transcript, error)
}

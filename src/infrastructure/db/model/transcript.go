package model

type TranscriptModel struct {
	ID     int64  `db:"id"`
	Text   string `db:"text"`
	NoteID int64  `db:"note_id"`
}

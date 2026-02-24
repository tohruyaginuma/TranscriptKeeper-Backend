package domain

type NoteRepository interface {
	Create(note *Note) (*Note, error)
	Update(note *Note) (*Note, error)
	Delete(id NoteID) error
	ListByUserID(userID UserID) ([]*Note, error)
}

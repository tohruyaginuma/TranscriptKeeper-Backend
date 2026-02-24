package application

type NoteService struct {
	noteRepository NoteRepository
}

func NewNoteService(noteRepository NoteRepository) *NoteService {
	return &NoteService{noteRepository: noteRepository}
}

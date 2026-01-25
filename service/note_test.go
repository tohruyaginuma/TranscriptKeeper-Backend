package service_test

// type fakeNoteRepo struct {
// 	createGotName domain.NoteName
// 	createID      domain.NoteID
// 	createErr     error

// 	listGotLimit  int
// 	listGotOffset int
// 	listNotes     []domain.Note
// 	listErr       error

// 	retrieveGotID domain.NoteID
// 	retrieveNote  domain.Note
// 	retrieveErr   error

// 	deleteGotID domain.NoteID
// 	deleteErr   error

// 	updateGotUser domain.Note
// 	updateID      domain.NoteID
// 	updateErr     error

// 	countResult int
// 	countErr    error
// }

// var _ sservice.NoteRepository = (*fakeNoteRepo)(nil)

// func (r *fakeNoteRepo) Create(ctx context.Context, userID domain.UserID, name domain.NoteName) (domain.NoteID, error) {
// 	r.createGotName = name
// 	return domain.NoteID(r.createID), r.createErr
// }
// func (r *fakeNoteRepo) List(ctx context.Context, userID domain.UserID, limit, offset int) (notes []domain.Note, err error) {
// }
// func (r *fakeNoteRepo) Delete(ctx context.Context, noteID domain.NoteID) (err error) {}
// func (r *fakeNoteRepo) Update(ctx context.Context, note *domain.Note) (noteID domain.NoteID, err error) {
// }
// func (r *fakeNoteRepo) Count(ctx context.Context, userID domain.UserID) (count int, err error) {}

// func TestNote_Create_OK()              {}
// func TestNote_Create_InvalidArgument() {}
// func TestNote_Create_RepoErr()         {}
// func TestNote_List_OK()                {}
// func TestNote_List_CountErr()          {}
// func TestNote_List_RepoErr()           {}
// func TestNote_Delete_OK()              {}
// func TestNote_Delete_NotFound()        {}
// func TestNote_Delete_RepoErr()         {}
// func TestNote_Update_OK()              {}
// func TestNote_Update_InvalidArgument() {}
// func TestNote_Update_NotFound()        {}
// func TestNote_Update_RepoErr()         {}

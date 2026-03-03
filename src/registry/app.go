package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/note"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/user"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/handler"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/db/postgres"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/external"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
)

type Registry struct {
	Firebase *external.Firebase

	UserHandler *handler.UserHandler
	NoteHandler *handler.NoteHandler

	UserUsecase *user.UserUsecase
}

func NewRegistry(db *sqlx.DB) *Registry {
	validator := lib.NewValidator()

	firebaseApp, err := external.NewFirebaseApp()
	if err != nil {
		panic(err)
	}
	firebase, err := external.NewFirebase(firebaseApp)
	if err != nil {
		panic(err)
	}

	userRepository := postgres.NewUserRepository(db)
	userUseCase := user.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	noteRepository := postgres.NewNoteRepository(db)
	noteUsecase := note.NewNoteUsecase(noteRepository)
	noteHandler := handler.NewNoteHandler(noteUsecase, validator)

	return &Registry{
		Firebase:    firebase,
		UserHandler: userHandler,
		UserUsecase: userUseCase,
		NoteHandler: noteHandler,
	}
}

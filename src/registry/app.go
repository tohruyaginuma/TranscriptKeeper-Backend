package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/handler"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/repository"
)

type Registry struct {
	Firebase    *repository.Firebase
	UserHandler *handler.UserHandler
	UserUsecase *application.UserUsecase
	NoteHandler *handler.NoteHandler
}

func NewRegistry(db *sqlx.DB) *Registry {
	validator := lib.NewValidator()

	firebaseApp, err := repository.NewFirebaseApp()
	if err != nil {
		panic(err)
	}
	firebase, err := repository.NewFirebase(firebaseApp)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUseCase := application.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	noteRepository := repository.NewNoteRepository(db)
	noteUsecase := application.NewNoteUsecase(noteRepository)
	noteHandler := handler.NewNoteHandler(noteUsecase, validator)

	return &Registry{
		Firebase:    firebase,
		UserHandler: userHandler,
		UserUsecase: userUseCase,
		NoteHandler: noteHandler,
	}
}

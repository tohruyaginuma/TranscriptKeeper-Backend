package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/handler"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/repository"
)

type Registry struct {
	Firebase    *repository.Firebase
	UserHandler *handler.UserHandler
}

func NewRegistry(db *sqlx.DB) *Registry {
	// validator := config.NewValidator()

	firebaseApp, err := repository.NewFirebaseApp()
	if err != nil {
		panic(err)
	}
	firebase, err := repository.NewFirebase(firebaseApp)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUseCase := application.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	// noteRepository := repository.NewNoteRepository(db)
	// noteService := service.NewNoteService(noteRepository)
	// noteHandler := handler.NewNoteHandler(validator, noteService)

	// transcriptionRepository := repository.NewTranscriptionRepository(db)
	// transcriptionService := service.NewTranscriptionService(transcriptionRepository, noteRepository)
	// transcriptionHandler := handler.NewTranscriptionHandler(validator, transcriptionService)

	return &Registry{
		Firebase:    firebase,
		UserHandler: userHandler,
	}
}

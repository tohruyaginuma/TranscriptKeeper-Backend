package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/handler"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/repository"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/service"
)

type Registry struct {
	UserHandler          UserHandler
	NoteHandler          NoteHandler
	TranscriptionHandler TranscriptionHandler
}

func NewRegistry(db *sqlx.DB) *Registry {
	validator := config.NewValidator()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(validator, userService)

	noteRepository := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(noteRepository)
	noteHandler := handler.NewNoteHandler(validator, noteService)

	transcriptionRepository := repository.NewTranscriptionRepository(db)
	transcriptionService := service.NewTranscriptionService(transcriptionRepository, noteRepository)
	transcriptionHandler := handler.NewTranscriptionHandler(validator, transcriptionService)

	return &Registry{
		UserHandler:          userHandler,
		NoteHandler:          noteHandler,
		TranscriptionHandler: transcriptionHandler,
	}
}

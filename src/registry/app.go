package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
)

type Registry struct {
}

func NewRegistry(db *sqlx.DB) *Registry {
	validator := config.NewValidator()

	// userRepository := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepository)
	// userHandler := handler.NewUserHandler(validator, userService)

	// noteRepository := repository.NewNoteRepository(db)
	// noteService := service.NewNoteService(noteRepository)
	// noteHandler := handler.NewNoteHandler(validator, noteService)

	// transcriptionRepository := repository.NewTranscriptionRepository(db)
	// transcriptionService := service.NewTranscriptionService(transcriptionRepository, noteRepository)
	// transcriptionHandler := handler.NewTranscriptionHandler(validator, transcriptionService)

	return &Registry{}
}

package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/note"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/transcript"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/user"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	noteHandler "github.com/tohruyaginuma/TranscriptKeeper-Backend/src/handler/note"
	transcriptHandler "github.com/tohruyaginuma/TranscriptKeeper-Backend/src/handler/transcript"
	userHandler "github.com/tohruyaginuma/TranscriptKeeper-Backend/src/handler/user"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/db/postgres"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/external"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
)

type Registry struct {
	Firebase *external.Firebase

	UserHandler       *userHandler.UserHandler
	NoteHandler       *noteHandler.NoteHandler
	TranscriptHandler *transcriptHandler.TranscriptHandler
	UserUsecase       *user.UserUsecase
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
	userHandler := userHandler.NewUserHandler(userUseCase)

	noteRepository := postgres.NewNoteRepository(db)
	noteUsecase := note.NewNoteUsecase(noteRepository)
	noteHandler := noteHandler.NewNoteHandler(noteUsecase, validator)

	cloudflareClient := external.NewCloudflareWorkersAIClient(config.CFAPIToken, config.CFAccountID)
	transcriptRepository := postgres.NewTranscriptRepository(db)
	transcriptUsecase := transcript.NewTranscriptUsecase(cloudflareClient, transcriptRepository)
	transcriptHandler := transcriptHandler.NewTranscriptHandler(transcriptUsecase, validator)

	return &Registry{
		Firebase:          firebase,
		UserHandler:       userHandler,
		UserUsecase:       userUseCase,
		NoteHandler:       noteHandler,
		TranscriptHandler: transcriptHandler,
	}
}

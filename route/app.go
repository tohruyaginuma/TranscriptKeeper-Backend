package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/registry"
)

func SetRoute(e *echo.Echo, r *registry.Registry) {
	const version = "v1"

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]any{"result": "OK"})
	})

	userGroup := e.Group(version + "/users")

	userGroup.POST("/:id/notes/:noteID/transcriptions", r.TranscriptionHandler.Create)
	userGroup.GET("/:id/notes/:noteID/transcriptions", r.TranscriptionHandler.ListByNoteID)

	userGroup.POST("/:id/notes", r.NoteHandler.Create)
	userGroup.GET("/:id/notes", r.NoteHandler.List)
	userGroup.PUT("/:id/notes/:noteID", r.NoteHandler.Update)
	userGroup.DELETE("/:id/notes/:noteID", r.NoteHandler.Delete)

	userGroup.POST("", r.UserHandler.CreateUser)
	userGroup.GET("", r.UserHandler.List)
	userGroup.GET("/:id", r.UserHandler.Retrieve)
	userGroup.PUT("/:id", r.UserHandler.Update)
	userGroup.DELETE("/:id", r.UserHandler.Delete)
}

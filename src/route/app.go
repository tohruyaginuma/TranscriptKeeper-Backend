package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/middleware"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/registry"
)

func SetRoute(e *echo.Echo, r *registry.Registry) {
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]any{"result": "OK"})
	})

	firebaseAuthMiddleware := middleware.NewFirebaseAuthMiddleware(r.Firebase)
	userIDMiddleware := middleware.NewUserIDMiddleware(r.UserUsecase)

	v1Group := e.Group(config.Version)

	v1Group.Use(firebaseAuthMiddleware.Handle)
	v1Group.POST("/auth", r.UserHandler.Authenticate)

	noteGroup := v1Group.Group("/notes")
	noteGroup.Use(userIDMiddleware.Handle)
	noteGroup.GET("", r.NoteHandler.ListByUserID)
	noteGroup.POST("", r.NoteHandler.Create)
	noteGroup.PUT("/:note_id", r.NoteHandler.Update)
	noteGroup.DELETE("/:note_id", r.NoteHandler.Delete)
}

package note

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
)

type NoteHandler struct {
	noteUsecase noteUsecase
	validator   *lib.Valid
}

func NewNoteHandler(noteUsecase noteUsecase, validator *lib.Valid) *NoteHandler {
	return &NoteHandler{noteUsecase: noteUsecase, validator: validator}
}

func (h *NoteHandler) Create(c echo.Context) error {
	userID, ok := c.Get(config.UserIDKey).(domain.UserID)
	if !ok {
		slog.Error("NoteHandler.Create", "error", "missing userID")

		return c.JSON(http.StatusUnauthorized, map[string]any{
			"result":  "NG",
			"message": "missing userID",
		})
	}

	slog.Info("NoteHandler.Create", "userID", userID)

	var noteRequest noteCreateRequest
	if err := c.Bind(&noteRequest); err != nil {
		slog.Error("NoteHandler.Create", "error", "invalid request body", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid request body",
		})
	}

	if err := h.validator.IsValid(noteRequest); err != nil {
		slog.Error("NoteHandler.Create", "error", "invalid request", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid request",
		})
	}

	title, err := domain.NewNoteTitle(noteRequest.Title)
	if err != nil {
		slog.Error("NoteHandler.Create", "error", "invalid request", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid request",
		})
	}
	ctx := c.Request().Context()
	noteID, err := h.noteUsecase.Create(ctx, title, userID)
	if err != nil {
		slog.Error("NoteHandler.Create", "error", "failed to create note", "error", err)

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to create note",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result":  "OK",
		"note_id": noteID,
	})
}

func (h *NoteHandler) Update(c echo.Context) error {
	userID, ok := c.Get(config.UserIDKey).(domain.UserID)
	if !ok {
		slog.Error("NoteHandler.Update", "error", "missing userID")

		return c.JSON(http.StatusUnauthorized, map[string]any{
			"result":  "NG",
			"message": "missing userID",
		})
	}

	slog.Info("NoteHandler.Update", "userID", userID)

	var noteRequest noteUpdateRequest
	if err := c.Bind(&noteRequest); err != nil {
		slog.Error("NoteHandler.Update", "error", "invalid request body", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid request body",
		})
	}

	if err := h.validator.IsValid(noteRequest); err != nil {
		slog.Error("NoteHandler.Update", "error", "invalid request", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid request",
		})
	}

	title, err := domain.NewNoteTitle(noteRequest.Title)
	if err != nil {
		slog.Error("NoteHandler.Update", "error", "invalid request", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid request",
		})
	}

	noteIDStr := c.Param("note_id")
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		slog.Error("NoteHandler.Update", "error", "invalid note_id", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}
	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("NoteHandler.Update", "error", "invalid note_id", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}

	ctx := c.Request().Context()
	err = h.noteUsecase.Update(ctx, noteID, title)
	if err != nil {
		slog.Error("NoteHandler.Update", "error", "failed to update note", "error", err)

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to update note",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]any{
		"result": "OK",
	})
}

func (h *NoteHandler) Delete(c echo.Context) error {
	userID, ok := c.Get(config.UserIDKey).(domain.UserID)
	if !ok {
		slog.Error("NoteHandler.Delete", "error", "missing userID")

		return c.JSON(http.StatusUnauthorized, map[string]any{
			"result":  "NG",
			"message": "missing userID",
		})
	}

	slog.Info("NoteHandler.Delete", "userID", userID)

	noteIDStr := c.Param("note_id")
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		slog.Error("NoteHandler.Delete", "error", "invalid note_id", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}
	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("NoteHandler.Delete", "error", "invalid note_id", "error", err)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}

	ctx := c.Request().Context()
	err = h.noteUsecase.Delete(ctx, noteID)
	if err != nil {
		slog.Error("NoteHandler.Delete", "error", "failed to delete note", "error", err)

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to delete note",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]any{
		"result": "OK",
	})
}

func (h *NoteHandler) ListByUserID(c echo.Context) error {
	userID, ok := c.Get(config.UserIDKey).(domain.UserID)
	if !ok {
		slog.Error("NoteHandler.ListByUserID", "error", "missing userID")

		return c.JSON(http.StatusUnauthorized, map[string]any{
			"result":  "NG",
			"message": "missing userID",
		})
	}

	slog.Info("NoteHandler.ListByUserID", "userID", userID)

	ctx := c.Request().Context()
	notes, err := h.noteUsecase.ListByUserID(ctx, userID)
	if err != nil {
		slog.Error("NoteHandler.ListByUserID", "error", "failed to list notes", "error", err)

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to list notes",
		})
	}

	response := newNoteResponse(notes)

	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
		"notes":  response.Notes,
	})
}

package transcript

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	transcriptUsecase "github.com/tohruyaginuma/TranscriptKeeper-Backend/src/application/transcript"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
)

type TranscriptHandler struct {
	usecase   usecase
	validator *lib.Valid
}

func NewTranscriptHandler(usecase usecase, validator *lib.Valid) *TranscriptHandler {
	return &TranscriptHandler{usecase: usecase, validator: validator}
}

func (h *TranscriptHandler) Create(c echo.Context) error {
	_, ok := c.Get(config.UserIDKey).(domain.UserID)
	if !ok {
		slog.Error("TranscriptHandler.Create", "error", "missing userID")

		return c.JSON(http.StatusUnauthorized, map[string]any{
			"result":  "NG",
			"message": "missing userID",
		})
	}

	noteIDStr := c.Param("note_id")
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		slog.Error("TranscriptHandler.Create", "error", "invalid note_id", "noteIDStr", noteIDStr)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}

	slog.Info("TranscriptHandler.Create", "noteID", noteIDInt)

	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("TranscriptHandler.Create", "error", "invalid note_id", "noteIDInt", noteIDInt)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}

	audio, err := c.FormFile("audio")
	if err != nil {
		slog.Error("TranscriptHandler.Create", "error", "missing audio")

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "missing audio",
		})
	}

	language := c.FormValue("language")
	if language == "" {
		slog.Error("TranscriptHandler.Create", "error", "missing language")

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "missing language",
		})
	}

	audioFile, err := audio.Open()
	if err != nil {
		slog.Error("TranscriptHandler.Create", "error", "failed to open audio file", "error", err)

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to open audio file",
		})
	}

	defer audioFile.Close()

	ctx := c.Request().Context()
	transcriptResult, err := h.usecase.Create(ctx, audioFile, language, noteID)
	if err != nil {
		slog.Error("TranscriptHandler.Create", "error", "failed to create transcript", "error", err)

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to create transcript",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result":        "OK",
		"transcript_id": transcriptResult.TranscriptID,
	})
}

func (h *TranscriptHandler) RetrieveByNoteID(c echo.Context) error {
	noteIDStr := c.Param("note_id")
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		slog.Error("TranscriptHandler.RetrieveByNoteID", "error", "invalid note_id", "noteIDStr", noteIDStr)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}

	slog.Info("TranscriptHandler.RetrieveByNoteID", "noteID", noteIDInt)

	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("TranscriptHandler.RetrieveByNoteID", "error", "invalid note_id", "noteIDInt", noteIDInt)

		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid note_id",
		})
	}

	ctx := c.Request().Context()
	transcript, err := h.usecase.RetrieveByNoteID(ctx, noteID)
	if err != nil {
		slog.Error("TranscriptHandler.RetrieveByNoteID", "error", "failed to retrieve transcript", "error", err)

		if errors.Is(err, transcriptUsecase.ErrTranscriptNotFound) {
			return c.JSON(http.StatusNotFound, map[string]any{
				"result":  "NG",
				"message": "transcript not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to retrieve transcript",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
		"text":   transcript.Text,
	})
}

package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/request"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/response"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/service"
)

type transcriptionHandler struct {
	validator *config.Valid
	service   TranscriptionService
}

func NewTranscriptionHandler(valid *config.Valid, service TranscriptionService) *transcriptionHandler {
	return &transcriptionHandler{
		validator: valid,
		service:   service,
	}
}

func (h *transcriptionHandler) Create(c echo.Context) error {
	noteIDStr := c.Param("noteID")
	slog.Debug("transcription create", "noteID", noteIDStr)

	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}
	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("convert to domain failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	var req request.CreateTranscriptionRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("bind request failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	if err := h.validator.IsValid(req); err != nil {
		slog.Error("request validation failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	transcriptionID, err := h.service.Create(c.Request().Context(), noteID, req.Content)
	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			slog.Error("note not found", "err", err)
			return c.JSON(http.StatusNotFound, map[string]any{
				"result": "NG",
			})
		}

		slog.Error("create transcription failed", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result":          "OK",
		"transcriptionID": transcriptionID.Value(),
	})
}

func (h *transcriptionHandler) ListByNoteID(c echo.Context) error {
	noteIDStr := c.Param("noteID")
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		slog.Error("parse int failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}
	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("convert to domain failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit := parseLimit(limitStr)
	offset := parseOffset(offsetStr)

	transcriptions, err := h.service.ListByNoteID(c.Request().Context(), noteID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			slog.Error("note not found", "err", err)
			return c.JSON(http.StatusNotFound, map[string]any{
				"result": "NG",
			})
		}

		slog.Error("list transcription failed", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	var transcriptionResponse = make([]response.TranscriptionResponse, len(transcriptions))
	for i, t := range transcriptions {
		transcriptionResponse[i] = response.TranscriptionResponse{
			ID:      t.ID().Value(),
			Content: t.Content().String(),
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result":         "OK",
		"transcriptions": transcriptionResponse,
	})
}

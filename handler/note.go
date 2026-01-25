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

type noteHandler struct {
	validator *config.Valid
	service   NoteService
}

func NewNoteHandler(valid *config.Valid, service NoteService) *noteHandler {
	return &noteHandler{
		validator: valid,
		service:   service,
	}
}

func (h *noteHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	userIDStr := c.Param("id")
	userIDInt, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		slog.Error("failed parse int for user id", "err: ", err)
		return c.JSON(http.StatusOK, map[string]any{
			"result": "NG",
		})
	}
	userID, err := domain.NewUserID(userIDInt)
	if err != nil {
		slog.Error("domain new user id faild")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	var req request.CreateNoteRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("failed bind response struct", "err: ", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	if err := h.validator.IsValid(req); err != nil {
		slog.Error("validator", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{})
	}

	noteID, err := h.service.Create(ctx, userID, req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
		"noteId": noteID,
	})
}

func (h *noteHandler) List(c echo.Context) error {
	userIDStr := c.Param("id")

	slog.Debug("noteHandler.List()", "userID", userIDStr)

	userIDInt, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		slog.Error("userID parseInt failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	userID, err := domain.NewUserID(userIDInt)
	if err != nil {
		slog.Error("convert userID domain failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit := parseLimit(limitStr)
	offset := parseOffset(offsetStr)

	notes, count, err := h.service.List(c.Request().Context(), userID, limit, offset)
	if err != nil {
		slog.Error("list notes failed", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	notesResponse := make([]response.NoteResponse, len(notes))
	for i, n := range notes {
		notesResponse[i] = response.NoteResponse{
			ID:   n.ID().Value(),
			Name: n.Name().String(),
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
		"notes":  notesResponse,
		"limit":  limit,
		"offset": offset,
		"total":  count,
	})
}

func (h *noteHandler) Delete(c echo.Context) error {
	noteIDStr := c.Param("noteID")

	slog.Debug("note handler delete", "noteID", noteIDStr)

	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		slog.Error("noteID parseInt failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}
	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("noteID convert failed to domain", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	if err := h.service.Delete(c.Request().Context(), noteID); err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			slog.Warn("delete execution failed", "err", err)
			return c.JSON(http.StatusNotFound, map[string]any{
				"result": "NG",
			})
		}

		slog.Error("delete execution failed", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *noteHandler) Update(c echo.Context) error {
	userIDStr := c.Param("id")
	noteIDStr := c.Param("noteID")

	slog.Debug("note handler update", "userID", userIDStr, "noteID", noteIDStr)

	userIDInt, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		slog.Error("userID parseint failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		slog.Error("noteID parseint failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	userID, err := domain.NewUserID(userIDInt)
	if err != nil {
		slog.Error("convert failed userID domain", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}
	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		slog.Error("convert failed noteID domain", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	var req = request.CreateNoteRequest{}
	if err := c.Bind(&req); err != nil {
		slog.Error("request bind failed", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	if err = h.service.Update(c.Request().Context(), noteID, req.Name, userID); err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			slog.Error("target note isn't exist", "err", err)
			return c.JSON(http.StatusNotFound, map[string]any{
				"result": "NG",
			})
		}
		slog.Error("udpate note failed", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
	})
}

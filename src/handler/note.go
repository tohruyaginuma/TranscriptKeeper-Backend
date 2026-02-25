package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
)

type NoteHandler struct {
	noteUsecase noteUsecase
	userUsecase userUsecase
	validator   *validator.Validate
}

func NewNoteHandler(noteUsecase noteUsecase) *NoteHandler {
	return &NoteHandler{noteUsecase: noteUsecase}
}

func (h *NoteHandler) Create(c echo.Context) error {
	userID, ok := c.Get(config.UserIDKey).(int64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"result":  "NG",
			"message": "missing userID",
		})
	}

	fmt.Println("userID", userID)

	return c.JSON(http.StatusOK, map[string]any{
		"result":  "OK",
		"user_id": userID,
	})
}

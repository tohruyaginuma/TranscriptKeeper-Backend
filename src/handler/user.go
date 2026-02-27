package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
)

type UserHandler struct {
	userUsecase userUsecase
	validator   *lib.Valid
}

func NewUserHandler(userUsecase userUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) Authenticate(c echo.Context) error {
	firebaseUIDStr, ok := c.Get(config.FirebaseUIDKey).(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"result":  "NG",
			"message": "missing firebaseUID",
		})
	}

	firebaseUID, err := domain.NewFirebaseUID(firebaseUIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid firebaseUID",
		})
	}

	ctx := c.Request().Context()
	userID, isCreated, err := h.userUsecase.CreateOrFindByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result":  "NG",
			"message": "failed to create or find by firebase_uid",
		})
	}

	if isCreated {
		return c.JSON(http.StatusCreated, map[string]any{
			"result":  "OK",
			"user_id": userID,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result":     "OK",
		"user_id":    userID,
		"is_created": isCreated,
	})
}

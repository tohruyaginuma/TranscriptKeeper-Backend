package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type userUsecase interface {
	FindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (user *domain.User, err error)
}

type UserIDMiddleware struct {
	userUsecase userUsecase
}

func NewUserIDMiddleware(userUsecase userUsecase) *UserIDMiddleware {
	return &UserIDMiddleware{userUsecase: userUsecase}
}

func (m *UserIDMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		user, err := m.userUsecase.FindByFirebaseUID(ctx, firebaseUID)
		if err != nil {
			if errors.Is(err, domain.ErrUserNotFound) {
				return c.JSON(http.StatusNotFound, map[string]any{
					"result":  "NG",
					"message": "user not found",
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]any{
				"result":  "NG",
				"message": "failed to find by firebase_uid",
			})
		}

		c.Set(config.UserIDKey, user.ID)
		return next(c)
	}
}

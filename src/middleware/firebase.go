package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
)

type TokenVerifier interface {
	VerifyIDToken(ctx context.Context, idToken string) (string, error)
}

type FirebaseAuthMiddleware struct {
	verifier TokenVerifier
}

func NewFirebaseAuthMiddleware(verifier TokenVerifier) *FirebaseAuthMiddleware {
	return &FirebaseAuthMiddleware{verifier: verifier}
}

func (m *FirebaseAuthMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header")
		}

		token := strings.TrimSpace(parts[1])
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		ctx := c.Request().Context()

		uid, err := m.verifier.VerifyIDToken(ctx, token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		c.Set(config.FirebaseUIDKey, uid)
		return next(c)
	}
}

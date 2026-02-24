package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/registry"
)

func SetRoute(e *echo.Echo, r *registry.Registry) {
	const version = "v1"

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]any{"result": "OK"})
	})

	userGroup := e.Group(version + "/users")

	
}

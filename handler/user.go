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

type userHandler struct {
	validator *config.Valid
	service   UserService
}

func NewUserHandler(valid *config.Valid, service UserService) *userHandler {
	return &userHandler{
		validator: valid,
		service:   service,
	}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	slog.Debug("UserHandler.CreateUser")

	ctx := c.Request().Context()

	var req request.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("UserHandler.CreateUser: not able to Bind")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	if err := h.validator.IsValid(req); err != nil {
		slog.Error("UserHandler.CreateUser: Validation error")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	userID, err := h.service.Create(ctx, req.Name)
	if err != nil {
		slog.Error("UserHandler.CreateUser: service error: ", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"result":  "OK",
		"user_id": userID,
	})
}

func (h *userHandler) List(c echo.Context) error {
	slog.Debug("UserHandler.List")

	ctx := c.Request().Context()
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	limit := parseLimit(limitStr)
	offset := parseOffset(offsetStr)

	users, count, err := h.service.List(ctx, limit, offset)
	if err != nil {
		slog.Error("UserHandler.List: service error: ", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	usersResponse := make([]response.UserResponse, len(users))
	for i, v := range users {
		usersResponse[i] = response.UserResponse{
			ID:   int64(v.ID()),
			Name: v.Name().String(),
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
		"users":  usersResponse,
		"limit":  limit,
		"offset": offset,
		"total":  count,
	})
}

func (h *userHandler) Retrieve(c echo.Context) error {
	slog.Debug("UserHandler.Retrieve")

	userIDStr := c.Param("id")
	userIDInt, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	userID, err := domain.NewUserID(userIDInt)
	if err != nil {
		slog.Error("userHandler.Retrieve: domain.NewUserID error", "err", err, "userID: ", userID)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	user, err := h.service.Retrieve(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]any{
				"result": "NG",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
		"user": response.UserResponse{
			ID:   int64(user.ID()),
			Name: user.Name().String(),
		},
	})
}

func (h *userHandler) Delete(c echo.Context) error {
	slog.Debug("UserHandler.Delete")

	userIDStr := c.Param("id")
	userIDInt, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	userID, err := domain.NewUserID(userIDInt)
	if err != nil {
		slog.Error("userHandler.Delete: domain.NewUserID error", "err", err, "userID: ", userID)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	err = h.service.Delete(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]any{
				"result": "NG",
			})
		}
		slog.Error("UserHandler.Delete: service error", "err: ", err, "userID: ", userID)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *userHandler) Update(c echo.Context) error {
	slog.Debug("userHandler.Update")

	userIDStr := c.Param("id")
	userIDInt, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		slog.Error("userHandler.Update: parse error", "id", userIDStr, "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	var req request.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		slog.Error("userHandler.Update: Bind error", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	if err := h.validator.IsValid(req); err != nil {
		slog.Error("userHandler.Update: Validation error", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	userID, err := domain.NewUserID(userIDInt)
	if err != nil {
		slog.Error("userHandler.Update: domain.NewUserID error", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result": "NG",
		})
	}

	if err = h.service.Update(c.Request().Context(), userID, req.Name); err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			slog.Warn("userHandler.Update: user not found", "err", err)

			return c.JSON(http.StatusNotFound, map[string]any{
				"result": "NG",
			})
		}
		if errors.Is(err, service.ErrInvalidArgument) {
			slog.Warn("userHandler.Update: user name not valid", "err", err)
			return c.JSON(http.StatusBadRequest, map[string]any{
				"result": "NG",
			})
		}

		slog.Error("userHandler.Update: Service error", "err", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"result": "NG",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
		"userID": userID,
	})
}

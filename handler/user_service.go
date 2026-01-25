package handler

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type UserService interface {
	Create(ctx context.Context, name string) (domain.UserID, error)
	List(ctx context.Context, limit, offset int) (users []domain.User, count int, err error)
	Retrieve(ctx context.Context, userID domain.UserID) (domain.User, error)
	Delete(ctx context.Context, userID domain.UserID) error
	Update(ctx context.Context, userID domain.UserID, name string) error
}

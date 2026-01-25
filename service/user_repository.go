package service

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type UserRepository interface {
	Create(ctx context.Context, name domain.UserName) (domain.UserID, error)
	Update(ctx context.Context, user domain.User) (domain.UserID, error)
	List(ctx context.Context, limit, offset int) ([]domain.User, error)
	Retrieve(ctx context.Context, userID domain.UserID) (domain.User, error)
	Delete(ctx context.Context, userID domain.UserID) error
	Count(ctx context.Context) (count int, err error)
}

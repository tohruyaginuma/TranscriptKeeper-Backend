package application

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type UserRepository interface {
	Create(ctx context.Context, googleID domain.GoogleID) (userID domain.UserID, err error)
	FindByGoogleID(ctx context.Context, googleID domain.GoogleID) (userID domain.UserID, err error)
	ExistsByGoogleID(ctx context.Context, googleID domain.GoogleID) (exists bool, err error)
}

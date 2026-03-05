package user

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type userUsecase interface {
	CreateOrFindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, isCreated bool, err error)
	FindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (user *domain.User, err error)
}

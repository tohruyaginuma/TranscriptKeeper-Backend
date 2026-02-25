package application

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type userRepository interface {
	Create(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, err error)
	FindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, err error)
}

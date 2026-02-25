package handler

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type userUsecase interface {
	CreateOrFindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, isCreated bool, err error)
}

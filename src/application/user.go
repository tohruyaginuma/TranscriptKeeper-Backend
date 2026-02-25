package application

import (
	"context"
	"errors"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserUseCase struct {
	userRepository userRepository
}

func NewUserUseCase(userRepository userRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) CreateOrFindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, isCreated bool, err error) {
	userID, err = u.userRepository.Create(ctx, firebaseUID)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateFirebaseUID) {
			userID, err = u.userRepository.FindByFirebaseUID(ctx, firebaseUID)
			if err != nil {
				if errors.Is(err, domain.ErrUserNotFound) {
					return 0, false, ErrUserNotFound
				}

				return 0, false, err
			}

			return userID, false, nil
		}

		return 0, false, err
	}

	return userID, true, nil
}

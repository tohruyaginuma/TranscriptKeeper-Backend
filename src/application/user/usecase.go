package user

import (
	"context"
	"errors"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserUsecase struct {
	userRepository userRepository
}

func NewUserUsecase(userRepository userRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) CreateOrFindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, isCreated bool, err error) {
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

func (u *UserUsecase) FindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (user *domain.User, err error) {
	userID, err := u.userRepository.FindByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}

	user = domain.NewUser(userID, firebaseUID)

	return user, nil
}

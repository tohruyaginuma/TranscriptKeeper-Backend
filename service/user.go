package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/repository"
)

var (
	ErrUserNotFound    = errors.New("User Not Found")
	ErrInvalidArgument = errors.New("Invalid Argument")
)

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, name string) (domain.UserID, error) {
	userName, err := domain.NewUserName(name)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidArgument, err)
	}

	userID, err := s.repo.Create(ctx, userName)
	if err != nil {
		return 0, fmt.Errorf("userService.Create repo.Create Failed: %w", err)
	}

	return userID, nil
}

func (s *userService) List(ctx context.Context, limit, offset int) (users []domain.User, count int, err error) {
	users, err = s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("UserService.List: %w", err)
	}

	count, err = s.repo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users")
	}

	return users, count, nil
}

func (s *userService) Retrieve(ctx context.Context, userID domain.UserID) (domain.User, error) {
	user, err := s.repo.Retrieve(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return domain.User{}, ErrUserNotFound
		}

		return domain.User{}, fmt.Errorf("userService.Retrieve userID=%d: %w", int64(userID), err)
	}

	return user, nil
}

func (s *userService) Delete(ctx context.Context, userID domain.UserID) error {
	if err := s.repo.Delete(ctx, userID); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		}

		return fmt.Errorf("userService.Delete userID=%d: %w", int64(userID), err)
	}

	return nil
}

func (s *userService) Update(ctx context.Context, userID domain.UserID, name string) error {
	userName, err := domain.NewUserName(name)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidArgument, err)
	}
	user := domain.NewUser(userID, userName)

	_, err = s.repo.Update(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		}

		return fmt.Errorf("userService.Update Repository error userID:%d, %w", int64(userID), err)
	}

	return nil
}

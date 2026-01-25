package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/handler"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/repository"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/service"
)

type fakeUserRepo struct {
	createGotName domain.UserName
	createID      domain.UserID
	createErr     error

	listGotLimit  int
	listGotOffset int
	listUsers     []domain.User
	listErr       error

	retrieveGotID domain.UserID
	retrieveUser  domain.User
	retrieveErr   error

	deleteGotID domain.UserID
	deleteErr   error

	updateGotUser domain.User
	updateID      domain.UserID
	updateErr     error

	countResult int
	countErr    error
}

var _ service.UserRepository = (*fakeUserRepo)(nil)

func (f *fakeUserRepo) Create(ctx context.Context, name domain.UserName) (domain.UserID, error) {
	f.createGotName = name
	return f.createID, f.createErr
}

func (f *fakeUserRepo) List(ctx context.Context, limit, offset int) ([]domain.User, error) {
	f.listGotLimit = limit
	f.listGotOffset = offset
	return f.listUsers, f.listErr
}

func (f *fakeUserRepo) Retrieve(ctx context.Context, userID domain.UserID) (domain.User, error) {
	f.retrieveGotID = userID
	return f.retrieveUser, f.retrieveErr
}
func (f *fakeUserRepo) Delete(ctx context.Context, userID domain.UserID) error {
	f.deleteGotID = userID
	return f.deleteErr
}
func (f *fakeUserRepo) Update(ctx context.Context, user domain.User) (domain.UserID, error) {
	f.updateGotUser = user
	return f.updateID, f.updateErr
}

func (f *fakeUserRepo) Count(ctx context.Context) (count int, err error) {
	return f.countResult, f.countErr
}

func setup(t *testing.T) (*fakeUserRepo, handler.UserService) {
	t.Helper()
	repo := &fakeUserRepo{}
	svc := service.NewUserService(repo)

	return repo, svc
}

func TestUser_Create_OK(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)
	repo.createID = 10

	gotID, err := svc.Create(context.Background(), "  Alice  ")
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if gotID != 10 {
		t.Fatalf("id = %d want 10", gotID)
	}
	if repo.createGotName.String() != "Alice" {
		t.Fatalf("repo got name = %q, want %q", repo.createGotName.String(), "Alice")
	}
}

func TestUser_Create_InvalidArgument(t *testing.T) {
	t.Parallel()

	_, svc := setup(t)
	_, err := svc.Create(context.Background(), "")
	if err == nil {
		t.Fatalf("Create() expected error, got nil")
	}
	if !errors.Is(err, service.ErrInvalidArgument) {
		t.Fatalf("Create() error = %v, want ErrInvalidArgument", err)
	}
}

func TestUser_List_OK(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)
	repo.listUsers = make([]domain.User, 2)
	repo.countResult = 10

	users, count, err := svc.List(context.Background(), 1, 0)
	if err != nil {
		t.Fatalf("List() err = %v", err)
	}
	if repo.listGotLimit != 1 {
		t.Fatalf("List() repo.listGotLimit = %d, want 1", repo.listGotLimit)
	}
	if repo.listGotOffset != 0 {
		t.Fatalf("List() repo.listGotOffset = %d, want 0", repo.listGotOffset)
	}
	if len(users) != 2 {
		t.Fatalf("List() users len = %d, want 2", len(users))
	}
	if count != 10 {
		t.Fatalf("List() count mismatch count = %d, want 10 ", count)
	}
}

func TestUser_List_RepoErr(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)

	repoErr := errors.New("db failed")
	repo.listErr = repoErr

	_, _, err := svc.List(context.Background(), 1, 0)
	if err == nil {
		t.Fatalf("List() expected error, got nil")
	}

	if !errors.Is(err, repoErr) {
		t.Fatalf("List() err = %v, want to wrap %v", err, repoErr)
	}
}

func TestUser_List_CountErr(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)

	repoErr := errors.New("db failed")
	repo.countErr = repoErr

	_, _, err := svc.List(context.Background(), 1, 0)
	if err == nil {
		t.Fatalf("List() expected error, got nil")
	}

	if !errors.Is(err, repoErr) {
		t.Fatalf("List() err = %v, want to wrap %v", err, repoErr)
	}
}

func TestUser_Retrieve_NotFound(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)
	repo.retrieveErr = repository.ErrUserNotFound

	userID, err := domain.NewUserID(1)
	if err != nil {
		t.Fatalf("NewUserID() err = %v", err)
	}

	_, err = svc.Retrieve(context.Background(), userID)
	if err == nil {
		t.Fatalf("Retrieve() expected error, got nil")
	}
	if !errors.Is(err, service.ErrUserNotFound) {
		t.Fatalf("Retrieve() err = %v, want = %v", err, service.ErrUserNotFound)
	}
}

func TestUser_Delete_NotFound(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)
	repo.deleteErr = repository.ErrUserNotFound

	userID, err := domain.NewUserID(1)
	if err != nil {
		t.Fatalf("NewUserID() err = %v", err)
	}

	err = svc.Delete(context.Background(), userID)
	if err == nil {
		t.Fatalf("Delete() expected error, got nil")
	}
	if !errors.Is(err, service.ErrUserNotFound) {
		t.Fatalf("Delete() err = %v, want = %v", err, service.ErrUserNotFound)
	}
}

func TestUser_Update_OK(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)
	userID, err := domain.NewUserID(1)
	if err != nil {
		t.Fatalf("NewUserID() err = %v", err)
	}

	err = svc.Update(context.Background(), userID, "  Tohru  ")
	if err != nil {
		t.Fatalf("Update() err = %v", err)
	}
	if repo.updateGotUser.ID() != userID {
		t.Fatalf("repo got id = %v, want %v", repo.updateGotUser.ID(), userID)
	}
	if repo.updateGotUser.Name().String() != "Tohru" {
		t.Fatalf("repo got name = %q, want %q", repo.updateGotUser.Name().String(), "Tohru")
	}
}

func TestUser_Update_InvalidArgument(t *testing.T) {
	t.Parallel()

	_, svc := setup(t)
	userID, err := domain.NewUserID(1)
	if err != nil {
		t.Fatalf("NewUserID() err = %v", err)
	}

	err = svc.Update(context.Background(), userID, "  ")
	if err == nil {
		t.Fatalf("Update() expect err, got nil")
	}
	if !errors.Is(err, service.ErrInvalidArgument) {
		t.Fatalf("Update() err= %v, want %v", err, service.ErrInvalidArgument)
	}
}

func TestUser_Update_NotFound(t *testing.T) {
	t.Parallel()

	repo, svc := setup(t)

	repo.updateErr = repository.ErrUserNotFound

	userID, err := domain.NewUserID(1)
	if err != nil {
		t.Fatalf("NewUserID() err = %v", err)
	}

	err = svc.Update(context.Background(), userID, "Tohru")
	if err == nil {
		t.Fatalf("Update() expect err, got nil")
	}
	if !errors.Is(err, service.ErrUserNotFound) {
		t.Fatalf("Update() err= %v, want %v", err, service.ErrUserNotFound)
	}
}

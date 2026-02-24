package repository

import (
	"database/sql"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) (*domain.User, error)                {}
func (r *UserRepository) FindByGoogleID(googleID domain.GoogleID) (*domain.User, error) {}
func (r *UserRepository) ExistsByGoogleID(googleID domain.GoogleID) (bool, error)       {}

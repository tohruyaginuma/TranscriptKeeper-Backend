package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, googleID domain.GoogleID) (userID domain.UserID, err error) {
	const query = `
		INSERT INTO users (google_id)
		VALUES ($1)
		RETURNING id
	;`

	var id int64
	err = r.db.GetContext(ctx, &id, query, googleID)
	if err != nil {
		return 0, err
	}

	userID, err = domain.NewUserID(id)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepository) FindByGoogleID(ctx context.Context, googleID domain.GoogleID) (userID domain.UserID, err error) {
	const query = `
		SELECT id
		FROM users 
		WHERE google_id = $1
	;`

	var id int64
	err = r.db.GetContext(ctx, &id, query, googleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrUserNotFound
		}
		return 0, err
	}

	userID, err = domain.NewUserID(id)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepository) ExistsByGoogleID(ctx context.Context, googleID domain.GoogleID) (exists bool, err error) {
	const query = `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE google_id = $1
		)
	;`

	err = r.db.GetContext(ctx, &exists, query, googleID)
	if err != nil {
		return false, err
	}

	return exists, nil
}

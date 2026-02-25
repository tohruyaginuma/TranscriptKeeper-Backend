package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, err error) {
	const query = `
		INSERT INTO users (firebase_uid)
		VALUES ($1)
		ON CONFLICT (firebase_uid) DO NOTHING 
		RETURNING id
	;`

	var id int64
	err = r.db.GetContext(ctx, &id, query, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, domain.ErrDuplicateFirebaseUID
		}
		return 0, err
	}

	userID, err = domain.NewUserID(id)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepository) FindByFirebaseUID(ctx context.Context, firebaseUID domain.FirebaseUID) (userID domain.UserID, err error) {
	const query = `
		SELECT id
		FROM users 
		WHERE firebase_uid = $1
	;`

	var id int64
	err = r.db.GetContext(ctx, &id, query, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, domain.ErrUserNotFound
		}
		return 0, err
	}

	userID, err = domain.NewUserID(id)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

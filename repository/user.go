package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, name domain.UserName) (domain.UserID, error) {
	const baseQuery = `
		INSERT INTO users ( name )
		VALUES ($1)
		RETURNING id
	;`

	var id int64
	strName := name.String()

	if err := u.db.QueryRowxContext(ctx, baseQuery, strName).Scan(&id); err != nil {
		return 0, err
	}

	userID := domain.UserID(id)
	return userID, nil
}

func (u *userRepository) List(ctx context.Context, limit, offset int) ([]domain.User, error) {
	const baseQuery = `
	SELECT 
		u.id, 
		u.name
	FROM users AS u
	ORDER BY u.id ASC
	LIMIT $1 OFFSET $2
	;`

	var userModels []userModel

	if err := u.db.SelectContext(ctx, &userModels, baseQuery, limit, offset); err != nil {
		return nil, err
	}

	users := make([]domain.User, len(userModels))
	for i, v := range userModels {
		user, err := v.toDomain()
		if err != nil {
			return nil, err
		}

		users[i] = user
	}

	return users, nil
}

func (u *userRepository) Retrieve(ctx context.Context, userID domain.UserID) (domain.User, error) {
	const baseQuery = `
		SELECT
			u.id,
			u.name
		FROM users AS u
		WHERE u.id = $1
		LIMIT 1
	;`

	m := userModel{}
	if err := u.db.GetContext(ctx, &m, baseQuery, int64(userID)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("UserRepository.Retrieve:user id=%d: %w", userID, err)
	}

	user, err := m.toDomain()
	if err != nil {
		return domain.User{}, fmt.Errorf("userModel.toDomain: %w", err)
	}

	return user, nil
}

func (u *userRepository) Delete(ctx context.Context, userID domain.UserID) error {
	const baseQuery = `
		DELETE FROM users WHERE id = $1
	;`

	res, err := u.db.ExecContext(ctx, baseQuery, int64(userID))
	if err != nil {
		return fmt.Errorf("userRepository.Delete userID=%d: %w", int64(userID), err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("userRepository.Delete rowsAffected userID=%d: %w", int64(userID), err)
	}
	if n == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, user domain.User) (domain.UserID, error) {
	const baseQuery = `
		UPDATE users AS u
		SET name = $2
		WHERE u.id = $1
		RETURNING u.id
	;`

	var userID int64
	if err := u.db.QueryRowContext(
		ctx,
		baseQuery,
		int64(user.ID()),
		user.Name().String(),
	).Scan(&userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.ID(), ErrUserNotFound
		}
		slog.Error("userRepository.Update: query error", "userID", int64(user.ID()), "err", err)
		return user.ID(), fmt.Errorf("userRepository.Update userID:%d, %w", int64(user.ID()), err)
	}

	return domain.UserID(userID), nil
}

func (u *userRepository) Count(ctx context.Context) (count int, err error) {
	const query = `
		SELECT COUNT(*)
		FROM users AS u
	;`

	if err = u.db.GetContext(ctx, &count, query); err != nil {
		return 0, fmt.Errorf("faild to count users: %w", err)
	}

	return count, nil
}

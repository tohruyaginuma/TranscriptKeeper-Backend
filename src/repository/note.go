package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
)

type NoteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) Create(ctx context.Context, title domain.NoteTitle, userID domain.UserID) (noteID domain.NoteID, err error) {
	const query = `
		INSERT INTO notes (title, user_id)
		VALUES ($1, $2)
		RETURNING id
	;`

	var id int64
	err = r.db.GetContext(ctx, &id, query, title, userID)
	if err != nil {
		return 0, err
	}

	noteID, err = domain.NewNoteID(id)
	if err != nil {
		return 0, err
	}

	return noteID, nil
}

func (r *NoteRepository) Update(ctx context.Context, noteID domain.NoteID, title domain.NoteTitle) (err error) {
	const query = `
		UPDATE notes
		SET title = $1
		WHERE id = $2
	;`

	res, err := r.db.ExecContext(ctx, query, title, noteID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return domain.ErrNoteNotFound
	}

	return nil
}

func (r *NoteRepository) Delete(ctx context.Context, noteID domain.NoteID) (err error) {
	const query = `
		DELETE FROM notes
		WHERE id = $1
	;`

	res, err := r.db.ExecContext(ctx, query, noteID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return domain.ErrNoteNotFound
	}
	return nil
}

func (r *NoteRepository) ListByUserID(ctx context.Context, userID domain.UserID) (notes []*domain.Note, err error) {
	const query = `
		SELECT id, title, updated_at
		FROM notes
		WHERE user_id = $1
		ORDER BY updated_at DESC
	;`
	var models []noteModel
	err = r.db.SelectContext(ctx, &models, query, userID)
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		noteID, err := domain.NewNoteID(model.ID)
		if err != nil {
			return nil, err
		}
		noteTitle, err := domain.NewNoteTitle(model.Title)
		if err != nil {
			return nil, err
		}
		note := domain.NewNote(noteID, noteTitle, userID)
		notes = append(notes, note)
	}

	return notes, nil
}

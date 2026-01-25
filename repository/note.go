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
	ErrNoteNotFound = errors.New("note not found")
)

type noteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) *noteRepository {
	return &noteRepository{
		db: db,
	}
}

func (r *noteRepository) Create(ctx context.Context, userID domain.UserID, name domain.NoteName) (domain.NoteID, error) {
	const query = `
		INSERT INTO notes (name, user_id)
		VALUES ($1, $2)
		RETURNING id
	;`

	var noteIDInt int64

	if err := r.db.QueryRowContext(ctx, query, name, userID).Scan(&noteIDInt); err != nil {
		return domain.NoteID(0), fmt.Errorf("create note: %w", err)
	}

	noteID, err := domain.NewNoteID(noteIDInt)
	if err != nil {
		return domain.NoteID(0), fmt.Errorf("new note: invalid id : %w", err)
	}

	return noteID, nil
}

func (r *noteRepository) List(ctx context.Context, userID domain.UserID, limit, offset int) (notes []domain.Note, err error) {
	const query = `
		SELECT 
			n.id, 
			n.name, 
			n.user_id
		FROM notes AS n
		WHERE n.user_id = $1
		ORDER BY n.id
		LIMIT $2 OFFSET $3
	`
	var noteModels []NoteModel
	if err = r.db.SelectContext(ctx, &noteModels, query, userID.Value(), limit, offset); err != nil {
		return notes, fmt.Errorf("sql execution error: %w", err)
	}
	notes = make([]domain.Note, len(noteModels))
	for i, m := range noteModels {
		note, err := m.toDomain()
		if err != nil {
			return notes, fmt.Errorf("convert note domain error: %w", err)
		}
		notes[i] = *note
	}

	return notes, nil
}

func (r *noteRepository) Delete(ctx context.Context, noteID domain.NoteID) (err error) {
	const query = `
		DELETE FROM notes AS n
		WHERE n.id = $1
	;`

	res, err := r.db.ExecContext(ctx, query, noteID.Value())
	if err != nil {
		return fmt.Errorf("delete note failed: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rowsAffected failed: %w", err)
	}
	if n == 0 {
		return fmt.Errorf("note not found: %w", ErrNoteNotFound)
	}

	return nil
}

func (r *noteRepository) Update(ctx context.Context, note *domain.Note) (noteID domain.NoteID, err error) {
	const query = `
		UPDATE notes
		SET name = $1
		WHERE id = $2
		RETURNING id
	;`

	var noteIDInt int64
	if err = r.db.GetContext(ctx, &noteIDInt, query, note.Name().String(), note.ID().Value()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNoteNotFound
		}
		return 0, fmt.Errorf("failed update note: %w", err)
	}

	slog.Debug("update debug",
		"err", err,
		"noteIDInt", noteIDInt,
		"inputID", note.ID().Value(),
		"inputName", note.Name().String(),
	)
	noteID, err = domain.NewNoteID(noteIDInt)
	if err != nil {
		return 0, fmt.Errorf("failed to convert noteID: %w", err)
	}
	return noteID, nil
}

func (r *noteRepository) Count(ctx context.Context, userID domain.UserID) (count int, err error) {
	const query = `
		SELECT COUNT(*)
		FROM notes AS n
		WHERE n.user_id = $1
	;`

	if err = r.db.GetContext(ctx, &count, query, userID.Value()); err != nil {
		return 0, fmt.Errorf("count failed: %w", err)
	}

	return count, nil
}

func (r *noteRepository) Exists(ctx context.Context, noteID domain.NoteID) (exists bool, err error) {
	const query = `
		SELECT EXISTS(
			SELECT 1 
			FROM notes
			WHERE id = $1
		)
	;`

	if err = r.db.GetContext(ctx, &exists, query, noteID.Value()); err != nil {
		return false, fmt.Errorf("exists failed: %w", err)
	}

	return exists, nil
}

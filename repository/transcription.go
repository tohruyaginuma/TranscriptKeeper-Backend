package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/domain"
)

type transcriptionRepository struct {
	db *sqlx.DB
}

func NewTranscriptionRepository(db *sqlx.DB) *transcriptionRepository {
	return &transcriptionRepository{
		db: db,
	}
}

func (r *transcriptionRepository) Create(ctx context.Context, noteID domain.NoteID, content domain.TranscriptionContent) (transcriptionID domain.TranscriptionID, err error) {
	const query = `
		INSERT INTO transcriptions (content, note_id)
		VALUES ($1, $2)
		RETURNING id
	;`

	var transcriptionIDInt int64
	if err = r.db.GetContext(ctx, &transcriptionIDInt, query, content.String(), noteID.Value()); err != nil {

		return domain.TranscriptionID(0), fmt.Errorf("failed to create transcription: %w", err)
	}

	transcriptionID, err = domain.NewTranscriptionID(transcriptionIDInt)
	if err != nil {
		return domain.TranscriptionID(0), err
	}

	return transcriptionID, nil
}

func (r *transcriptionRepository) ListByNoteID(ctx context.Context, noteID domain.NoteID, limit, offset int) (transcriptions []domain.Transcription, err error) {
	const query = `
		SELECT 
			t.id,
			t.content,
			t.note_id
		FROM transcriptions AS t
		WHERE t.note_id = $1
		ORDER BY t.id DESC
		LIMIT $2 OFFSET $3
	;`

	var transcriptionModels []transcriptionModel
	if err := r.db.SelectContext(ctx, &transcriptionModels, query, noteID.Value(), limit, offset); err != nil {
		return nil, fmt.Errorf("list by note failed: %w", err)
	}

	transcriptions = make([]domain.Transcription, len(transcriptionModels))
	for i, t := range transcriptionModels {
		transcriptions[i], err = t.toDomain()
		if err != nil {
			return nil, fmt.Errorf("failed convert to domain: %w", err)
		}
	}

	return transcriptions, nil
}

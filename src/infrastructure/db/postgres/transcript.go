package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/domain"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/db/model"
)

type TranscriptRepository struct {
	db *sqlx.DB
}

func NewTranscriptRepository(db *sqlx.DB) *TranscriptRepository {
	return &TranscriptRepository{db: db}
}

func (r *TranscriptRepository) Create(ctx context.Context, text string, noteID domain.NoteID) (transcriptID domain.TranscriptID, err error) {
	const query = `
		INSERT INTO transcripts (text, note_id)
		VALUES ($1, $2)
		RETURNING id
	;`

	var id int64
	err = r.db.GetContext(ctx, &id, query, text, noteID)
	if err != nil {
		return 0, err
	}

	transcriptID, err = domain.NewTranscriptID(id)
	if err != nil {
		return 0, err
	}

	return transcriptID, nil
}

func (r *TranscriptRepository) RetrieveByNoteID(ctx context.Context, noteID domain.NoteID) (transcript domain.Transcript, err error) {
	const query = `
		SELECT id, text, note_id
		FROM transcripts
		WHERE note_id = $1
	;`

	var model model.TranscriptModel
	err = r.db.GetContext(ctx, &model, query, noteID)
	if err != nil {
		return domain.Transcript{}, err
	}

	transcriptID, err := domain.NewTranscriptID(model.ID)
	if err != nil {
		return domain.Transcript{}, err
	}

	noteID, err = domain.NewNoteID(model.NoteID)
	if err != nil {
		return domain.Transcript{}, err
	}

	transcript = *domain.NewTranscript(transcriptID, model.Text, noteID)

	return transcript, nil
}

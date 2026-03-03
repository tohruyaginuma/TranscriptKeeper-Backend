package repository

import (
	"github.com/jmoiron/sqlx"
)

type TranscriptRepository struct {
	db *sqlx.DB
}

func NewTranscriptRepository(db *sqlx.DB) *TranscriptRepository {
	return &TranscriptRepository{db: db}
}

package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
)

type TranscriptHandler struct {
	transcriptUsecase transcriptUsecase
	validator         *lib.Valid
}

func NewTranscriptHandler(transcriptUsecase transcriptUsecase, validator *lib.Valid) *TranscriptHandler {
	return &TranscriptHandler{transcriptUsecase: transcriptUsecase, validator: validator}
}

func (h *TranscriptHandler) Create(c echo.Context) error {
	return nil
}

func (h *TranscriptHandler) ListByNoteID(c echo.Context) error {
	return nil
}

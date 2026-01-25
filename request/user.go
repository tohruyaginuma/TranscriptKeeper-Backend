package request

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateNoteRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateNoteRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateTranscriptionRequest struct {
	Content string `json:"content" validate:"required"`
}

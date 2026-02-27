package handler

type noteCreateRequest struct {
	Title string `json:"title" validate:"required"`
}

type noteUpdateRequest struct {
	Title string `json:"title" validate:"required"`
}

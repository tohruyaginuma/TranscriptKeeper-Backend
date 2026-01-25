package response

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type NoteResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

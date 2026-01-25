package registry

import "github.com/labstack/echo/v4"

type UserHandler interface {
	CreateUser(c echo.Context) error
	List(c echo.Context) error
	Retrieve(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type NoteHandler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type TranscriptionHandler interface {
	Create(c echo.Context) error
	ListByNoteID(c echo.Context) error
}

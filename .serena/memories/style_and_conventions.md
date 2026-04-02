# Code Style and Conventions

## Language: Go 1.25

## Naming
- **Packages**: lowercase, short, singular (e.g., `note`, `user`, `postgres`)
- **Types**: PascalCase (e.g., `NoteHandler`, `NoteUsecase`, `NoteID`)
- **Functions/Methods**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase
- **Constants**: PascalCase or SCREAMING_SNAKE_CASE for config keys
- **Errors**: `var Err<Name> = errors.New(...)` pattern in domain layer

## Value Objects (Domain Layer)
- Domain types wrap primitives: `type NoteID int64`, `type NoteTitle string`
- Constructor functions validate: `func NewNoteID(id int64) (NoteID, error)`
- Validation errors defined as package-level vars: `var ErrInvalidNoteID = errors.New(...)`

## Error Handling
- Errors returned from all functions (no panics except startup)
- JSON error responses use `map[string]any{"result": "NG", "message": "..."}` format
- Successful responses use `map[string]any{"result": "OK", ...}`

## Logging
- `log/slog` throughout with structured key-value pairs
- Pattern: `slog.Error("HandlerName.MethodName", "error", "description", "error", err)`
- Pattern: `slog.Info("HandlerName.MethodName", "key", value)`

## Dependency Injection
- Interfaces defined in `port.go` per package (handler/application layers)
- Registry (`src/registry/app.go`) wires all dependencies
- Constructor injection pattern (`NewXxx(dep Dep) *Xxx`)

## HTTP Handlers
- Use Echo v4 context
- Extract `userID` from Echo context via `c.Get(config.UserIDKey).(domain.UserID)`
- Bind request body, validate with `lib.Valid`, then construct domain values
- HTTP status: 200 OK, 204 NoContent, 400 BadRequest, 401 Unauthorized, 500 InternalServerError

## No Docstrings / Comments
- Code is written without docstrings or inline comments unless logic is non-obvious

## Tests
- Table-driven tests in `_test.go` files alongside source
- Tests mainly in `src/domain/` layer
- No mocking of DB (integration-style)

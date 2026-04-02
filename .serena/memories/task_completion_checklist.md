# Task Completion Checklist

When completing a coding task in this project, do the following:

1. **Format**: `gofmt -w .` or let the editor handle it
2. **Vet**: `go vet ./...` — check for common mistakes
3. **Test**: `go test ./...` — ensure all tests pass
4. **Build check**: `go build ./src` — confirm it compiles

## Notes
- Tests are mainly in `src/domain/` — add domain tests for new domain logic
- No dedicated linter config found (no `.golangci.yml`) — `go vet` is the minimum
- If adding new dependencies: `go mod tidy`

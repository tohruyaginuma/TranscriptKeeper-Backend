# Suggested Commands

## Run
```bash
go run ./src          # Start the server on :8080
```

## Live Reload (air)
```bash
air                   # Hot reload using .air.toml config
```

## Test
```bash
go test ./...         # Run all tests (mainly domain layer)
go test ./src/domain/...  # Run domain tests only
```

## Build
```bash
go build ./src        # Build the binary
```

## Lint / Format
```bash
gofmt -w .            # Format Go code (standard tool)
go vet ./...          # Static analysis
```

## Database Setup
```bash
psql -h localhost -U transcript_keeper_local -d transcript_keeper_local -f ddl/init.sql
```

## Docker (local postgres)
```bash
docker-compose up -d  # Start local PostgreSQL via docker-compose
```

## Dependencies
```bash
go mod tidy           # Tidy go.mod / go.sum
go mod download       # Download dependencies
```

## Utilities (Darwin/macOS)
- `ls`, `find`, `grep`, `cat`, `sed`, `awk` — standard macOS/BSD variants
- `git` for version control

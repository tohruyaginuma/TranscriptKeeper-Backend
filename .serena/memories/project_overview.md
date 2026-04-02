# TranscriptKeeper Backend - Project Overview

## Purpose
Backend API for Transcript Keeper app. Authenticates users via Firebase, stores notes/transcripts in PostgreSQL, and uses Cloudflare Workers AI (whisper-large-v3-turbo) for audio transcription.

## Tech Stack
- **Language**: Go 1.25
- **Web Framework**: Echo v4
- **Database**: PostgreSQL (via sqlx + lib/pq)
- **Auth**: Firebase Admin SDK
- **External AI**: Cloudflare Workers AI
- **Validation**: go-playground/validator v10
- **Live Reload**: air (.air.toml)
- **Module**: `github.com/tohruyaginuma/TranscriptKeeper-Backend`

## Architecture (Layered)
```
src/
├── main.go
├── handler/         # HTTP handlers (note, transcript, user)
├── middleware/      # Firebase auth, user ID extraction, CORS
├── application/     # Use cases (note, transcript, user)
├── domain/          # Domain models and validation (value objects, errors)
├── infrastructure/
│   ├── db/postgres/ # PostgreSQL repositories
│   ├── db/model/    # DB model structs
│   └── external/    # Cloudflare, Firebase clients
├── registry/        # Dependency injection wiring
├── route/           # Route definitions
├── config/          # Config loading, logger, echo setup, constants
└── lib/             # Validator wrapper, utilities
```

## Key Relationships
- `users -> notes -> transcripts` (DB hierarchy)
- All `/v1/*` endpoints require `Authorization: Bearer <Firebase ID Token>`
- `/v1/auth` must be called once to create the app user before using note APIs
- Server runs on port 8080 (fixed)

## Environment Variables
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSL_MODE`
- `CLIENT_URL_WEB`, `CLIENT_URL_DESKTOP` (CORS origins)
- `GOOGLE_APPLICATION_CREDENTIALS` (full service account JSON string, not file path)
- `CF_API_TOKEN`, `CF_ACCOUNT_ID`

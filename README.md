# Transcript Keeper Backend

This is the backend API for Transcript Keeper.  
It uses Firebase Authentication for user verification, PostgreSQL for storing notes and transcripts, and Cloudflare Workers AI for audio transcription.

## Overview

- Authenticates users with Firebase ID tokens
- Creates an application user automatically on first authentication
- Supports creating, listing, updating, and deleting notes
- Accepts audio uploads and generates transcripts
- Stores and retrieves transcripts by note

## Tech Stack

- Go 1.25
- Echo
- PostgreSQL
- Firebase Admin SDK
- Cloudflare Workers AI (`@cf/openai/whisper-large-v3-turbo`)

## Architecture

The project follows a loosely layered structure.

- `src/handler`: HTTP handlers
- `src/middleware`: authentication, user resolution, and CORS
- `src/application`: use cases
- `src/domain`: domain models and validation
- `src/infrastructure`: PostgreSQL and external service integrations
- `src/registry`: dependency wiring
- `src/route`: route definitions

## Requirements

- Go 1.25 or later
- PostgreSQL 16 recommended
- A Firebase project with service account credentials
- A Cloudflare account with access to Workers AI

## Environment Variables

The application primarily uses the following environment variables.

| Name | Required | Default | Description |
| --- | --- | --- | --- |
| `DB_HOST` | No | `localhost` | PostgreSQL host |
| `DB_PORT` | No | `5432` | PostgreSQL port |
| `DB_USER` | No | `test_local` | PostgreSQL user |
| `DB_PASSWORD` | No | `password` | PostgreSQL password |
| `DB_NAME` | No | `test_local` | PostgreSQL database name |
| `DB_SSL_MODE` | No | `disable` | PostgreSQL SSL mode |
| `CLIENT_URL_WEB` | No | `http://localhost:5173` | Allowed CORS origin for the web client |
| `CLIENT_URL_DESKTOP` | No | `http://localhost:5174` | Allowed CORS origin for the desktop client |
| `GOOGLE_APPLICATION_CREDENTIALS` | Yes | - | Firebase Admin SDK service account JSON string |
| `CF_API_TOKEN` | Yes | - | Cloudflare API token |
| `CF_ACCOUNT_ID` | Yes | - | Cloudflare account ID |

Notes:

- `GOOGLE_APPLICATION_CREDENTIALS` is expected to contain the full service account JSON as a string, not a file path.
- The server port is currently fixed to `8080`.

## Local Setup

### 1. Prepare PostgreSQL

Start PostgreSQL using your preferred method and create a database for the application.

Example:

```sql
CREATE DATABASE transcript_keeper_local;
CREATE USER transcript_keeper_local WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE transcript_keeper_local TO transcript_keeper_local;
```

### 2. Apply the schema

```bash
psql -h localhost -U transcript_keeper_local -d transcript_keeper_local -f ddl/init.sql
```

### 3. Configure environment variables

Use `.envrc.example` as a base, or export the variables directly in your shell.

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=transcript_keeper_local
export DB_PASSWORD=password
export DB_NAME=transcript_keeper_local
export DB_SSL_MODE=disable
export CLIENT_URL_WEB=http://localhost:5173
export CLIENT_URL_DESKTOP=http://localhost:5174
export GOOGLE_APPLICATION_CREDENTIALS='{"type":"service_account","project_id":"..."}'
export CF_API_TOKEN=your_cloudflare_api_token
export CF_ACCOUNT_ID=your_cloudflare_account_id
```

### 4. Start the application

```bash
go run ./src
```

The server will start on `http://localhost:8080`.

## API

All `/v1/*` endpoints require `Authorization: Bearer <Firebase ID Token>`.

### Health Check

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/` | Health check endpoint |

### Auth

| Method | Path | Description |
| --- | --- | --- |
| `POST` | `/v1/auth` | Creates or fetches the application user associated with the Firebase UID |

### Notes

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/v1/notes` | Returns the authenticated user's notes |
| `POST` | `/v1/notes` | Creates a new note |
| `GET` | `/v1/notes/:note_id` | Returns note details |
| `PUT` | `/v1/notes/:note_id` | Updates the note title |
| `DELETE` | `/v1/notes/:note_id` | Deletes the note |

### Transcripts

| Method | Path | Description |
| --- | --- | --- |
| `POST` | `/v1/notes/:note_id/transcripts` | Uploads audio and creates a transcript |
| `GET` | `/v1/notes/:note_id/transcripts` | Returns the transcript associated with the note |

## Request Examples

### Authenticate

```bash
curl -X POST http://localhost:8080/v1/auth \
  -H "Authorization: Bearer <FIREBASE_ID_TOKEN>"
```

### Create Note

```bash
curl -X POST http://localhost:8080/v1/notes \
  -H "Authorization: Bearer <FIREBASE_ID_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Meeting Memo"}'
```

### Upload Audio and Create a Transcript

```bash
curl -X POST http://localhost:8080/v1/notes/1/transcripts \
  -H "Authorization: Bearer <FIREBASE_ID_TOKEN>" \
  -F "audio=@sample.m4a" \
  -F "language=ja"
```

### Get a Transcript

```bash
curl -X GET http://localhost:8080/v1/notes/1/transcripts \
  -H "Authorization: Bearer <FIREBASE_ID_TOKEN>"
```

## Database Schema

`ddl/init.sql` defines the following three tables.

- `users`
- `notes`
- `transcripts`

The relationship is `users -> notes -> transcripts`.

## Testing

```bash
go test ./...
```

At the moment, tests are mainly included for the domain layer.

## Project Structure

```text
.
├── ddl
├── src
│   ├── application
│   ├── config
│   ├── domain
│   ├── handler
│   ├── infrastructure
│   ├── middleware
│   ├── registry
│   └── route
├── docker-compose.yml
├── go.mod
└── README.md
```

## Notes

- Allowed CORS origins are controlled by `CLIENT_URL_WEB` and `CLIENT_URL_DESKTOP`.
- Audio files are base64-encoded before being sent to Cloudflare Workers AI.
- Even after Firebase authentication succeeds, `/v1/auth` must be called once before using the note APIs so the application user can be created.

# CodeBridge Backend

This is the Go backend for CodeBridge, structured according to domain-driven design and idiomatic Go practices.

## Architecture

- **`cmd/server/`**: Application entrypoint (`main.go`).
- **`internal/`**: Domain modules containing handlers (HTTP/WS transport), services (business logic & state machines), repositories (PostgreSQL persistence), models, and DTOs.
  - `auth/`: JWT, sessions, password hashing.
  - `users/`: User profile management.
  - `projects/`: Project lifecycle & state machine.
  - `github/`: GitHub API client & repo synchronization.
  - `matching/`: Freelancer matching algorithms.
  - `workspace/`: Docker sandbox orchestration.
  - `ai/`: RAG pipeline & LLM integrations.
  - `submissions/`: Submission & review tracking.
  - `payments/`: Escrow simulation & payment state machine.
  - `notifications/`: WebSocket & event dispatching.
  - `admin/`: Admin metrics & moderation.
  - `middleware/`: Auth, RBAC, logging, recovery.
  - `database/`: PostgreSQL connection & migration manager.
  - `config/`: Environment configuration loader.
  - `shared/`: Errors, responses, validation, logger, utils.
- **`migrations/`**: SQL migration files.
- **`seeds/`**: Seed data scripts.
- **`tests/`**: Unit, integration, e2e, and security tests.
- **`docker/`**: Dockerfile & docker-compose for backend services.
- **`docs/`**: API specifications, architecture diagrams, and DB schemas.

## State Machines

1. **Project**: `Draft` → `Open` → `Assigned` → `In Progress` → `Submitted` → `Changes Requested / Approved` → `Completed`
2. **Workspace**: `Requested` → `Creating` → `Ready` → `Active` → `Stopped` → `Archived`
3. **Submission**: `Draft` → `Submitted` → `Under Review` → `Changes Requested / Approved`
4. **Payment**: `Unrequested` → `Authorized` → `Held` → `Release Pending` → `Released`

## API Base Route

Base URL: `/api/v1`

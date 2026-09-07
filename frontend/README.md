# CodeBridge Frontend

This is the React frontend client for CodeBridge.

## Architecture

- **`src/pages/`**: Top-level page views (Auth, Dashboard, Projects, Freelancers, Workspace, Review, Payments, Profile, Admin).
- **`src/components/`**: Modular, reusable UI components organized by domain.
- **`src/services/`**: API and WebSocket communication layer interacting with backend endpoints.
- **`src/hooks/`**: Custom React hooks.
- **`src/context/`**: React context providers.
- **`src/store/`**: Client-side state stores.
- **`src/types/`**: JSDoc/TypeScript-compatible type definitions.
- **`src/routes/`**: Client-side routing definitions.
- **`src/utils/`**: Helper utilities and formatters.
- **`src/constants/`**: Frontend constant declarations.
- **`tests/`**: Unit, integration, and E2E UI test suites.

## Communication

- Connects to Backend REST API at `/api/v1`
- Connects to Backend WebSocket server for real-time notifications and workspace state updates.
- No backend business rules or direct DB connections in frontend code.

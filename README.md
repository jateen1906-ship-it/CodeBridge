# CodeBridge 🌉

CodeBridge is an automated collaborative platform connecting clients and verified freelancers with integrated GitHub workspaces, AI/RAG analysis, Dockerized sandboxes, and milestone-based escrow payments.

## Monorepo Architecture

This repository is organized as a clean monorepo for a 2-person team:

- **`backend/`** (Developer 1 - Backend): Go modular/domain-driven backend, PostgreSQL, WebSockets, Docker, AI/RAG, RBAC.
- **`frontend/`** (Developer 2 - Frontend): React UI, Vite, state management, WebSocket client, REST API integration.

```
codebridge/
├── backend/       # Go backend workspace
├── frontend/      # React frontend workspace
├── .gitignore     # Root git ignore
└── README.md      # Project overview & workspace setup
```

## Team Workspaces & Responsibilities

| Role | Workspace | Tech Stack | Primary Focus |
|------|-----------|------------|---------------|
| **Backend Developer** | `/backend` | Go, PostgreSQL, Docker, WebSockets | REST API, WebSockets, Auth/RBAC, GitHub API, AI/RAG, Workspaces, Escrow |
| **Frontend Developer** | `/frontend` | React, Vite, WebSockets | UI/UX, Pages, Components, Client State, API/WS Consumer |

## Getting Started

Refer to the respective README files:
- [Backend Setup & Architecture](./backend/README.md)
- [Frontend Setup & Architecture](./frontend/README.md)

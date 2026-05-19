## Architecture Principles

This project follows a separation-of-concerns approach.

### Frontend

The React frontend is organized by features. UI components are responsible for presentation and user interaction, while API calls and business-specific utilities are separated into service or feature-specific files.

### Backend

The Go backend is organized around business domains such as users, goals, exercises, meals, and nutrition.

Each backend domain may contain:

- `handler.go` for HTTP request and response handling
- `service.go` for business logic
- `repository.go` for database access contracts
- `model.go` for domain models and data structures

This structure keeps the code readable and aligned with common backend patterns used in professional software engineering teams.

## Backend Request Flow

```text
HTTP Request
    ↓
Handler
    ↓
Service
    ↓
Repository
    ↓
Database

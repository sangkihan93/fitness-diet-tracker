# ADR 0003: Use Feature-Based Backend Structure

## Status

Accepted

## Context

The backend includes several business domains, including users, goals, exercises, meals, and nutrition.

A purely technical folder structure could make it harder to understand which files belong to which business concept.

## Decision

Organize backend code by business domain. Each domain can contain a handler, service, repository, and model.

Example:

```text
internal/exercises/
├── handler.go
├── service.go
├── repository.go
└── model.go
```

## Consequences

### Positive

- Clear separation of concerns
- Easier to understand business objects
- Easier to test each domain
- Aligns with common controller/service/repository patterns

### Negative

- Some repeated file names across folders
- Requires consistent naming discipline

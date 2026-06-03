# ADR 0001: Use a Monorepo

## Status

Accepted

## Context

The project contains a React frontend, Go backend, infrastructure configuration, and documentation.

## Decision

Use a monorepo with separate application folders under `apps/`.

## Consequences

### Positive

- Easier portfolio review
- Single pull request workflow
- Shared documentation
- Easier local development

### Negative

- CI must be configured carefully by folder
- Project organization must remain clear

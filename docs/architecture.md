# Architecture

## Overview

Fitness Diet Tracker is a web-first, mobile-responsive application that allows users to track exercise, meals, nutrition, dietary restrictions, and progress toward fitness goals.

The project is built as a professional full-stack portfolio project using React, TypeScript, Go, AWS-ready architecture, and USWDS-inspired accessibility standards.

## High-Level Architecture

```text
React Web App
    |
    | HTTPS / JSON
    v
Go REST API
    |
    | Business logic
    v
Repository Layer
    |
    | Database queries
    v
Database
```

Authentication is planned through AWS Cognito.

## Frontend

The frontend is built with React and TypeScript.

Frontend responsibilities:

- Render UI
- Collect user input
- Display validation messages
- Call backend APIs
- Handle responsive layouts
- Follow USWDS-inspired accessibility patterns

The frontend should not contain complex backend business rules.

## Backend

The backend is built with Go.

Backend responsibilities:

- Handle HTTP requests
- Validate incoming data
- Apply business rules
- Manage user-specific resources
- Communicate with the database
- Validate authentication tokens eventually

## Backend Layering

Each business domain follows this pattern:

```text
handler.go
    ↓
service.go
    ↓
repository.go
    ↓
database
```

### Handler

The handler receives HTTP requests and returns HTTP responses.

It should not contain heavy business logic.

### Service

The service contains business rules.

Examples:

- Validate that an exercise duration is greater than zero
- Validate that daily calorie targets are reasonable
- Calculate nutrition progress
- Check whether a meal fits dietary restrictions

### Repository

The repository handles database access.

Examples:

- Save an exercise log
- Load meals for a user
- Find a user goal
- Update nutrition targets

### Model

The model defines domain objects.

Examples:

- UserProfile
- FitnessGoal
- ExerciseLog
- MealLog
- NutritionTarget
- DietaryRestriction

## Planned AWS Architecture

Potential AWS services:

- Amazon Cognito for authentication
- Amazon RDS PostgreSQL for relational data
- AWS App Runner, ECS Fargate, or Lambda/API Gateway for backend deployment
- AWS Amplify Hosting or S3 + CloudFront for frontend hosting
- CloudWatch for logs
- Secrets Manager for sensitive configuration

## Design Standards

The UI should follow USWDS-inspired principles:

- Accessible forms
- Clear labels
- Keyboard navigation
- Responsive design
- Strong color contrast
- Clear error states
- Simple content hierarchy

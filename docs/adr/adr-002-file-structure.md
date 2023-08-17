# ADR 002: File Structure for Video Management API

**Title**: File structure for video Management API

**Date**: 13/08/2023

**Status**: Accepted

## Context

The Video Management API project aims to provide a RESTful API for managing videos and related annotations. As the project grows, it's essential to maintain a clear and scalable directory structure that separates concerns and aligns with best practices.

## Decision

We have decided to adopt a file structure that emphasizes the separation of concerns, aligns with the principles of Domain-Driven Design (DDD) and hexagonal architecture, and facilitates modularity and maintainability.

The proposed file structure is:

```
/video-management-api/
|-- cmd/
|   |-- main.go                   # Entry point of the application
|
|-- app/
|   |-- /adapters/                # Services and Repository Implementations

|
|   |-- /api/
|   |   |-- handler.go            # HTTP handlers for the RESTful API
|   |   |-- routes.go             # API routes definition
|
|   |-- infra/
|   |   |-- auth/
|   |   |   |-- service.go        # Authentication service (JWT generation, validation)
|   |   |   |-- middleware.go     # Middleware for JWT verification
|   |
|   |   |-- config/
|   |   |   |-- config.go         # Configuration loading (e.g., from environment variables)
|   |
|   |
|   |   |-- db/
|   |   |   |-- migration.sql     # SQL migrations for initializing the database
|   |   |   |-- connection.go     # Database connection setup and utilities
|
|   |-- domain/
|   |   |-- ports/                # Services and Repository Interfaces
|   |   |-- model/
|   |   |   |-- user.go           # User struct and related methods
|   |   |   |-- video.go          # Video struct and related methods
|   |   |   |-- annotation.go     # Annotation struct and related methods
```

## Rationale

- **Separation of Concerns**: By dividing the project into `api`, `infra`, `domain`, and `test` directories, we ensure that each part of the application has a distinct responsibility.
  
- **Domain-Driven Design (DDD)**: The `domain` folder houses the core business logic and domain models, emphasizing the importance of the business domain in the project's architecture.
  
- **Hexagonal Architecture**: The `infra` folder encapsulates all infrastructure-related code, ensuring that the core domain logic remains decoupled from external concerns like authentication, database operations, and configuration.
  
- **Scalability and Maintainability**: This structure allows for easy expansion as the project grows. New features, services, or models can be added without disrupting the existing structure.

## Consequences

- Developers will have a clear understanding of where specific code resides, leading to faster development and debugging.
  
- The project will be more maintainable and scalable as it grows.
  
- The clear separation of concerns will reduce the risk of tightly coupled code, making future refactoring or feature additions smoother.

---

This ADR provides a clear rationale for the chosen file structure and outlines the benefits and consequences of the decision. It serves as a reference for current and future team members to understand the reasoning behind the project's organization.
# go-ddd-starter

A minimal Go starter project using Domain-Driven Design (DDD) principles, PostgreSQL, and a clean separation between configuration, domain logic, persistence, and HTTP handlers.

## Features

- DDD-inspired project layout
- PostgreSQL connection with `sqlx`
- Database migrations via `sql-migrate`
- Environment-driven configuration using `godotenv`
- Basic user signup and login endpoints
- JWT token creation for authentication
- Simple middleware stack for CORS, logging, and preflight handling

## Project Structure

- `main.go`: application entrypoint
- `cmd/serve.go`: bootstraps configuration, database, services, and HTTP server
- `config/`: environment configuration loader
- `infra/db/`: PostgreSQL connection and migration helpers
- `domain/`: business entities and domain models
- `internal/user/`: user service and repository interfaces
- `repo/`: concrete repository implementation for `User`
- `rest/`: HTTP server and middleware setup
- `rest/handlers/user/`: user-related request handling
- `utils/`: helper functions for JWT, password hashing, and responses
- `migrations/`: SQL migration files for database schema

## Prerequisites

- Go 1.20+ installed
- PostgreSQL database available
- `git` installed

## Setup

1. Copy or create a `.env` file in the project root.

```env
VERSION=0.1.0
SERVICE_NAME=go-ddd-starter
HTTP_PORT=8080
JWT_SECRET_KEY=your_jwt_secret_key

DB_HOST=localhost
DB_PORT=5432
DB_NAME=your_database
DB_USER=your_user
DB_PASSWORD=your_password
DB_ENABLE_SSL_MODE=false
```

2. Create the PostgreSQL database specified in `DB_NAME`.

3. Install Go module dependencies.

```bash
go mod download
```

## Run the Server

Start the application from the project root:

```bash
go run .
```

The server will automatically run migrations from `./migrations` and listen on the port defined by `HTTP_PORT`.

## API Endpoints

### Sign Up

- `POST /users/sign-up`
- Request body:

```json
{
  "email": "user@example.com",
  "password": "secret",
  "firstName": "Jane",
  "lastName": "Doe",
  "username": "janedoe",
  "role": "user"
}
```

### Login

- `POST /users/login`
- Request body:

```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

- Response includes an `accessToken` and user payload.

## Notes

- Passwords are hashed before storage.
- JWT tokens are created with the secret from `JWT_SECRET_KEY`.
- Database migrations are executed automatically on startup.

## Customize

Use this project as a starter for a new Go API by:

- adding new domain entities under `domain/`
- extending services in `internal/`
- creating new repository implementations in `repo/`
- adding handlers and routes under `rest/handlers/`

## License

This project is provided as a starter template.

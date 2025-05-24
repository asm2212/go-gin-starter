# Go Gin Starter Kit

A starter kit for building a web API in Go using Gin, with Docker, configuration, dependency injection, and PostgreSQL support.

## Features

- Gin web framework
- PostgreSQL integration
- Environment variable/config management
- Docker support
- Example handler/service

## Getting Started

### Prerequisites

- Go 1.21+
- Docker
- PostgreSQL

### Setup

1. **Clone the repo**
   ```bash
   git clone https://github.com/asm2212/go-gin-starter.git
   cd go-gin-starter
   ```

2. **Set up your `.env` file**
   ```
   cp .env.example .env
   # Edit .env as needed
   ```

3. **Run with Docker**
   ```
   docker build -t go-gin-starter .
   docker run --env-file .env -p 8080:8080 go-gin-starter
   ```

4. **Run locally**
   ```bash
   go run ./cmd/main.go
   ```

5. **Test the API**
   ```
   curl http://localhost:8080/hello
   ```

## Project Structure

- `cmd/` — Entrypoint
- `config/` — Config loader
- `internal/` — App code (handlers, services, db)
- `pkg/` — Dependency injection setup

## License

MIT
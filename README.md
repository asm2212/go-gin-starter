# Go Gin Starter Kit

A starter kit for building a web API in Go using Gin, with Docker, configuration, dependency injection, JWT authentication, PostgreSQL support, and user authentication out of the box.

## Features

- Gin web framework
- PostgreSQL integration (with migration example)
- Environment variable/config management
- JWT authentication (registration, login, protected routes)
- Docker support (build and run with ease)
- Example handler/service structure
- Clean architecture: separation of handlers, services, repositories
- Ready for production: .env config, error handling, extensible modules

## Getting Started

### Prerequisites

- Go 1.21+ ([How to install Go](https://go.dev/doc/install))
- Docker
- PostgreSQL

---

## Quickstart

### 1. Clone the Repository

```bash
git clone https://github.com/asm2212/go-gin-starter.git
cd go-gin-starter
```

### 2. Configure Environment

Copy and edit the `.env` file:

```bash
cp .env.example .env
# Edit .env to match your local or Docker PostgreSQL settings and set a JWT_SECRET
```

Example:

```
APP_PORT=8080
APP_ENV=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=starterdb
JWT_SECRET=your_super_secret_key
```

---

### 3. Set Up Database

Make sure PostgreSQL is running and a database named `starterdb` exists. To create with Docker:

```bash
docker run --name go-starter-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=starterdb -p 5432:5432 -d postgres
```

Run the migration SQL:

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);
```

---

### 4. Run the Server

#### Locally (with Go):

```bash
go mod tidy
go run ./cmd/main.go
```

#### With Docker:

```bash
docker build -t go-gin-starter .
docker run --env-file .env -p 8080:8080 go-gin-starter
```

Or use Docker Compose (recommended for public use, see below).

---

## Public Usage (Recommended: Docker Compose)

**Best for reproducible, containerized deployments.**

1. Create a `docker-compose.yml` (or use the provided one if exists):

```yaml
version: "3.8"
services:
  db:
    image: postgres:16
    restart: always
    environment:
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: starterdb
    ports:
      - "5432:5432"
    volumes:
      - db_data:/var/lib/postgresql/data

  api:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    depends_on:
      - db

volumes:
  db_data:
```

2. Launch both services:

```bash
docker-compose up --build
```

---

## API Usage

### Register

```http
POST /register
Content-Type: application/json

{
  "username": "youruser",
  "password": "yourpass"
}
```

### Login

```http
POST /login
Content-Type: application/json

{
  "username": "youruser",
  "password": "yourpass"
}
```
_Response:_ `{ "token": "<jwt>" }`

### Authenticated endpoint

```http
GET /api/profile
Authorization: Bearer <jwt>
```

---

## Expanding the Starter

- Add new models in `internal/model/`
- Add repositories in `internal/db/`
- Add business logic in `internal/service/`
- Add new endpoints in `internal/api/`
- Add middleware in `internal/middleware/`
- Update your database schema and migrations as needed

---

## Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

---

## License

MIT

---

## Author

[asm2212](https://github.com/asm2212)
# Bank Backend API

This is a robust and scalable backend service for a simple banking application, built with Go. It provides RESTful APIs for user management, accounts, and financial transactions.

## 🚀 Technologies Used

- **Language:** [Go (Golang)](https://go.dev/)
- **Web Framework:** [Gin](https://gin-gonic.com/)
- **Database:** PostgreSQL (with `pgx` driver)
- **Database Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **SQL Code Generation:** [sqlc](https://sqlc.dev/)
- **Authentication:** JWT and PASETO (Platform-Agnostic Security Tokens)
- **Configuration:** [Viper](https://github.com/spf13/viper)
- **Containerization:** Docker & Docker Compose
- **Testing:** `testing`, `testify`, and `mockgen`

## 🛠 Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://go.dev/doc/install) (1.20+ recommended)
- [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/) (optional, but recommended for running Makefile commands)
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) CLI (for running DB migrations locally)
- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) (for generating DB code)

## 🏃 Getting Started

### Using Docker (Recommended)

The easiest way to get the application running is via Docker Compose, which sets up both the PostgreSQL database and the API server.

```bash
# Start the database and API server in the background
docker-compose up -d

# View the logs
docker-compose logs -f
```

The API will be available at `http://localhost:8080`.

### Local Development Setup

If you prefer to run the application locally (without Docker for the API):

1. **Start the PostgreSQL database container:**
   ```bash
   make postgres
   ```

2. **Create the database:**
   ```bash
   make createdb
   ```

3. **Run database migrations:**
   ```bash
   make migrateup
   ```

4. **Start the HTTP server:**
   ```bash
   make server
   ```

## 🧪 Testing

To run the full suite of unit tests with coverage:

```bash
make test
```

## 📜 Useful Makefile Commands

- `make postgres`: Starts a PostgreSQL container
- `make createdb` / `make dropdb`: Creates/Drops the database
- `make migrateup` / `make migratedown`: Runs database migrations
- `make sqlc`: Generates Go CRUD code from SQL using `sqlc`
- `make mock`: Generates mock DB interfaces for testing
- `make server`: Starts the Go HTTP server

## 🔌 gRPC Support (Ongoing)

Currently in the process of adding gRPC support to the API for high-performance internal communication. You might notice some gRPC-related files (`pb/`, `proto/`, `gapi/`) and Makefile commands (`make proto`, `make evans`). This feature is still experimental and under active development.

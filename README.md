# Go GraphQL Address API

This project is a GraphQL API built with Go that serves address data. It was created as a learning exercise to explore GraphQL server development in Go, including database migration, testing, and authentication with API keys.

---

## ✨ Features

- **GraphQL API:** Exposes address data through a clean GraphQL interface.
- **SQLite Database:** Uses a local SQLite database for data persistence.
- **Database Migration:** Includes a script to migrate data from JSON files into the database.
- **Authentication:** Protects specific queries using an API key and GraphQL directives.
- **Repository Pattern:** Separates database logic from the API layer for clean architecture.
- **Integration & Unit Tests:** Includes tests for both the repository (with a real in-memory DB) and the resolvers (with mocks).
- **Dockerized:** Comes with a `Dockerfile` for easy containerization and deployment.

---

## 🛠️ Tech Stack

- **Language:** [Go](https://go.dev/)
- **GraphQL Server:** [gqlgen](https://gqlgen.com/)
- **Database:** [SQLite](https://www.sqlite.org/)
- **Containerization:** [Docker](https://www.docker.com/)
- **HTTP Middleware:** [rs/cors](https://github.com/rs/cors)

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.21 or later)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) (for containerizing)
- `make` (optional, for using the Makefile shortcuts)

### 1. Installation

Clone the repository to your local machine:

```bash
git clone git@github.com:xaaha/address-api.git
cd address-api

```

- Start the server

```bash

make server
```

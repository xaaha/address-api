# ---- Build Stage ----
FROM golang:1.25.1-alpine3.22 AS builder

RUN apk add --no-cache git sqlite

WORKDIR /app

# Download modules first (better caching)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build server binary
RUN CGO_ENABLED=0 go build -o /server ./cmd/server

# Generate SQLite DB
RUN go run ./cmd/setup/main.go

# ---- Runtime Stage ----

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache sqlite

# Copy from builder stage 
COPY --from=builder /server /server
COPY --from=builder /app/internal/db/data.db /data/data.db

EXPOSE 8080

# Environment variable 
ENV DB_FILE_PATH=/data/data.db

ENTRYPOINT ["/server"]

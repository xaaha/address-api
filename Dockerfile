FROM golang:1.25.1-alpine3.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application, creating a static binary.
# CGO_ENABLED=0 is critical for building with the Go SQLite driver in Alpine Linux.
# -o /server builds the binary and places it in the root directory named 'server'.
RUN CGO_ENABLED=0 go build -o /server ./cmd/server


FROM alpine:latest

COPY --from=builder /server /server

# Copy your pre-populated database file into the image.
# This path should match where your DB is located.
COPY ./internal/db/data.db /data/data.db

EXPOSE 8080

# Set environment variables for the container.
# The API_KEY will be set when we run the container. ????
ENV DB_PATH=/data/data.db

# The command to run when the container starts
ENTRYPOINT ["/server"]

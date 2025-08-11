migrate: 
	@go run cmd/setup/main.go

server:
	@go run cmd/server/main.go

# keep adding folders with test instead of blanket coverage ...
test:
	@go test ./internal/...


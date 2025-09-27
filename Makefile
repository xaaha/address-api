migrate: 
	@go run cmd/setup/main.go

server:
	@go run cmd/server/main.go

# keep adding folders with test instead of blanket coverage ...
test:
	@go test ./internal/... ./graph/

# generate the resolver code form the schema.graphqls file
generate:
	@go run github.com/99designs/gqlgen generate

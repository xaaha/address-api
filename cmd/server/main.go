// Package main in Server is for server work like db migration and stuff
package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xaaha/address-api/graph"
	"github.com/xaaha/address-api/internal/repository"
)

// TODO: Move these hardcoded values into a config struct that
// is populated from environment variables
const (
	dbFile      = "internal/db/data.db"
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		slog.Error("failed to open db", "error", err, "dbFile", dbFile)
		os.Exit(1)
	}
	defer db.Close()

	addressRepository := repository.NewAddressRepository(db)

	resolver := &graph.Resolver{Repo: addressRepository}
	gqlSrv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: resolver}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", gqlSrv)

	slog.Info("connect for GraphQL playground", "url", fmt.Sprintf("http://localhost:%s/", port))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		slog.Error("HTTP server failed", "port", port, "error", err)
		os.Exit(1)
	}
}

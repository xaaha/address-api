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
// is populated from environment variables. For these two, values it's fine as they are
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

	addressRepo := repository.NewAddressRepository(db)
	rootResolver := &graph.Resolver{Repo: addressRepo}

	// This is the `c` variable from the documentation.
	c := graph.Config{Resolvers: rootResolver}

	//  Manually assign  Auth function to the generated Directives struct.
	// gqlgen created `c.Directives.Auth` because we have `@auth` in schema.
	c.Directives.Auth = graph.Auth

	executableSchema := graph.NewExecutableSchema(c)

	gqlSrv := handler.NewDefaultServer(executableSchema)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", graph.AuthMiddleWare(gqlSrv))

	slog.Info("connect for GraphQL playground", "url", fmt.Sprintf("http://localhost:%s/", port))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		slog.Error("HTTP server failed", "port", port, "error", err)
		os.Exit(1)
	}
}

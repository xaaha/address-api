// Package main in Server is for server work like db migration and stuff
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xaaha/address-api/graph"
)

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
		log.Fatalf("error ")
	}
	defer db.Close()

	resolver := graph.Resolver{DB: db}
	gqlSrv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", gqlSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

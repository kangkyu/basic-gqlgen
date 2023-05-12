package main

import (
	"basic-gqlgen/graph"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cfg := graph.Config{Resolvers: graph.NewResolver()}
	execSchema := graph.NewExecutableSchema(cfg)

	// Create a new GraphQL server with the executable schema.
	srv := handler.NewDefaultServer(execSchema)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

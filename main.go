package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	gql "github.com/nirajgeorgian/gateway/src/gql"
	resolver "github.com/nirajgeorgian/gateway/src/gql/resolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	c := gql.Config{Resolvers: &resolver.Resolver{}}
	gqlHandler := handler.GraphQL(
		gql.NewExecutableSchema(c),
		handler.IntrospectionEnabled(true),
	)
	http.Handle("/query", gqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

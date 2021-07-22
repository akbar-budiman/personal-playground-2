package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/akbar-budiman/personal-playground-2/graph"
	"github.com/akbar-budiman/personal-playground-2/graph/generated"
	"github.com/akbar-budiman/personal-playground-2/service"
)

var (
	redisAddress = "localhost:6379"
	defaultPort  = "8090"
)

func main() {
	service.InitializeLocalRedisConnectionPool(redisAddress)

	service.InitializeLocalCrdbPool()
	service.InitializeLocalCrdbDdl()

	service.RegisterConsumer()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

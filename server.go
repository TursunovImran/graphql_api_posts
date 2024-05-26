package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TursunovImran/graphql_api_posts/graph"

	"github.com/joho/godotenv"

	"github.com/TursunovImran/graphql_api_posts/graph/database"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	database.ConnectDB()
	database.MigrateDB()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
																Resolvers: &graph.Resolver{
																	Database: database.DBInstance},
																}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

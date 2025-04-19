package main

import (
	"log"
	"net/http"

	"github.com/alimikegami/go-monorepo/graphql-server/config"
	"github.com/alimikegami/go-monorepo/graphql-server/db"
	"github.com/alimikegami/go-monorepo/graphql-server/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	config, _ := config.LoadConfig()

	db, _ := db.InitDB(config)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Db: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("connect to http://localhost:8081/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

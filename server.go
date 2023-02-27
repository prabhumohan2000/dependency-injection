package main

import (
	"log"
	"net/http"
	"os"

	"githu.com/prabhumohan2000/test_8/db"
	"githu.com/prabhumohan2000/test_8/graph"
	"githu.com/prabhumohan2000/test_8/users"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8000"

func main() {
	db.Init() // initialize db init function
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userService := &users.UserServicefile{
		DB: db.DbInstance,
	} // create instance of userStruct passing it to user Struct

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserService: userService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

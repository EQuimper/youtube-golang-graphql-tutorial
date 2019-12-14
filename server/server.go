package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v9"

	"github.com/equimper/meetmeup"
	"github.com/equimper/meetmeup/postgres"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "meetmeup_dev",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := meetmeup.Config{Resolvers: &meetmeup.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   postgres.UsersRepo{DB: DB},
	}}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(meetmeup.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TomTomRixRix/Feedbackbuch/server/graph"
	"github.com/TomTomRixRix/Feedbackbuch/server/graph/generated"
	"github.com/TomTomRixRix/Feedbackbuch/server/graph/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "testdb.sqlite3")
	if err != nil {
		os.Exit(1)
	}

	db.AutoMigrate(&model.Comment{})

	return db
}

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: initDB(), ToBeNotified: make([]graph.Notification, 0)}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

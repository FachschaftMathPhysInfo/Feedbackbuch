package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/FachschaftMathPhysInfo/Feedbackbuch/server/graph"
	"github.com/FachschaftMathPhysInfo/Feedbackbuch/server/graph/generated"
	"github.com/FachschaftMathPhysInfo/Feedbackbuch/server/graph/model"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

func initDB() *gorm.DB {
	databaseConnectionType := "postgres"
	databaseConnectionString := os.Getenv("DB_CONNECTION_STRING")

	if databaseConnectionString != "" {
		log.Println(databaseConnectionString)
	} else {
		log.Println("DB_CONNECTION_STRING empty!")
		os.Exit(1)
	}
	err := fmt.Errorf("initial connect failed")

	db, err := gorm.Open(databaseConnectionType, databaseConnectionString)
	for err != nil {
		db, err = gorm.Open(databaseConnectionType, databaseConnectionString)
		time.Sleep(500 * time.Millisecond)
		log.Println(err)
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

	router := chi.NewRouter()
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	es := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: initDB(), ToBeNotified: make([]graph.Notification, 0)}})
	srv := handler.New(es)

	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return true // r.Host == "localhost"
			}, ReadBufferSize: 1024,
			WriteBufferSize: 1024},
		KeepAlivePingInterval: 100 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	router.Handle("/query", srv)
	fmt.Print(
		"==========================================\n",
		"Started the backend listening on Port "+port+"\n",
		"==========================================\n",
	)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

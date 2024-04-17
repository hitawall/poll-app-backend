package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
	"poll-app-backend/ent"
	"poll-app-backend/handlers"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=poll_app sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := httprouter.New()
	router.POST("/login", handlers.LoginHandler(client))
	router.POST("/signup", handlers.SignupHandler(client))
	router.POST("/polls", handlers.CreatePoll(client))
	router.POST("/polls/:id/options", handlers.AddOption(client))
	router.POST("/options/:id/vote", handlers.VoteOption(client))
	router.GET("/polls", handlers.GetPolls(client))
	router.GET("/options/:id/voters", handlers.GetVoters(client))
	router.PUT("/options/:id", handlers.UpdateOption(client))

	// Setup CORS to allow specific origins and methods
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Only allow frontend to connect
		AllowCredentials: true,                              // Allow credentials
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"}, // You might need to adjust these headers depending on your application needs
		ExposedHeaders:   []string{"X-Custom-Header"},               // Example of exposing custom headers
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

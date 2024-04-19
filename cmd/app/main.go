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

	// Apply authentication middleware only to these routes
	router.POST("/polls", handlers.AuthMiddleware(client, handlers.CreatePoll(client)))
	router.GET("/polls/:id", handlers.AuthMiddleware(client, handlers.GetPoll(client)))
	router.POST("/polls/:id/options", handlers.AuthMiddleware(client, handlers.AddOption(client)))
	router.POST("/options/:id/vote", handlers.AuthMiddleware(client, handlers.VoteOption(client)))
	router.GET("/polls", handlers.AuthMiddleware(client, handlers.GetPolls(client)))
	router.GET("/options/:id/voters", handlers.AuthMiddleware(client, handlers.GetVoters(client)))
	router.PUT("/options/:id", handlers.AuthMiddleware(client, handlers.UpdateOption(client)))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

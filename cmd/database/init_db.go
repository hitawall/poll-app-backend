package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 // Default PostgreSQL port
	user     = "postgres"
	password = "password"
	dbname   = "poll_app" // Specific database for the application
)

func main() {
	// Connection string for the default database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)

	// Connecting to PostgreSQL using the default database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Ensuring connection is established
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping PostgreSQL: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	// Check if the specific database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	// Create the database if it does not exist
	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		fmt.Printf("Database %s created successfully\n", dbname)
	} else {
		fmt.Printf("Database %s already exists\n", dbname)
	}
}

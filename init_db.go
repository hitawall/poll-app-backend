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
	dbname   = "postgres" // Default database to connect to before creating a new one
)

func init_db() {
	// Setting up the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connecting to PostgreSQL
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensuring connection is established
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	// Create a new database
	_, err = db.Exec("CREATE DATABASE poll_app")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database created successfully")

	// Connect to the new database to create tables
	psqlInfoNewDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=poll_app sslmode=disable",
		host, port, user, password)
	dbNew, err := sql.Open("postgres", psqlInfoNewDB)
	if err != nil {
		log.Fatal(err)
	}
	defer dbNew.Close()

	// Create tables
	// Replace this SQL command with your actual table creation command
	_, err = dbNew.Exec("CREATE TABLE IF NOT EXISTS users(email TEXT NOT NULL PRIMARY KEY, name TEXT NOT NULL, password TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")
}

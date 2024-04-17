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
	_, err = dbNew.Exec("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, email VARCHAR(255) UNIQUE NOT NULL, name VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users Table created successfully")

	// Replace this SQL command with your actual table creation command
	_, err = dbNew.Exec("CREATE TABLE IF NOT EXISTS polls(id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, creator_id INTEGER NOT NULL, FOREIGN KEY (creator_id) REFERENCES users(id))")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Polls Table created successfully")

	// Replace this SQL command with your actual table creation command
	_, err = dbNew.Exec("CREATE TABLE IF NOT EXISTS polloptions(id SERIAL PRIMARY KEY, text VARCHAR(255) NOT NULL, votes INTEGER DEFAULT 0, poll_id INTEGER NOT NULL, FOREIGN KEY (poll_id) REFERENCES polls(id))")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Poll Options Table created successfully")

	// Replace this SQL command with your actual table creation command
	_, err = dbNew.Exec("CREATE TABLE IF NOT EXISTS votes(id SERIAL PRIMARY KEY, voted_on TIMESTAMP NOT NULL, user_id INTEGER NOT NULL, polloption_id INTEGER NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id), FOREIGN KEY (polloption_id) REFERENCES polloptions(id), UNIQUE (user_id, polloption_id))")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Votes Table created successfully")
}

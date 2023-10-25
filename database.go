package main

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func ConnectToDatabase(serverName string, connectionString string) {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("Error connecting to db")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to create db connection:" + err.Error())
	}

	DB = db
	fmt.Printf("Connected to database server %s\n", serverName)
}

func InitializeDatabaseConnections() map[string]*sql.DB {
	databaseConnections := make(map[string]*sql.DB)

	// TODO: Init all the db connection here. Look into how I can get the value from users.
	return databaseConnections
}

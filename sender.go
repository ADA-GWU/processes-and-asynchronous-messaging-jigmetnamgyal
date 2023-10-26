package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var senderName = "YourName"

func sendMessageToServer(serverName string, db *sql.DB, message string) {
	// TODO
}

func main() {
	// TODO make the port dynamic
	databaseConnections := InitializeDatabaseConnections()

	var message string
	fmt.Print("Enter a message: ")
	_, err := fmt.Scan(&message)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for serverName, db := range databaseConnections {
		wg.Add(1)
		go func(serverName string, db *sql.DB) {
			defer wg.Done()
			sendMessageToServer(serverName, db, message)
		}(serverName, db)
	}

	wg.Wait()
}

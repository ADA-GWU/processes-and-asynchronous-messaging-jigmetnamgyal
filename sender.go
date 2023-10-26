package main

import (
	"database/sql"
	"fmt"
	"log"
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
}

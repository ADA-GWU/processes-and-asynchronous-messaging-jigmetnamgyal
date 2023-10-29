package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"processes-and-asynchronous-messaging-jigmetnamgyal/db"
	"sync"
	"time"
)

var senderName = "YourName"

func sendMessageToServer(serverName string, db *sql.DB, message string) {
	query := `
		INSERT INTO ASYNC_MESSAGE (SENDER_NAME, MESSAGE, "CURRENT_TIME") VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, senderName, message, time.Now())
	if err != nil {
		fmt.Println(err)
		log.Printf("Error sending message to %s: %v", serverName, err)
	}
}

func main() {
	databaseConnections := db.InitializeDatabaseConnections()

	var message string
	fmt.Print("Enter a message: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	} else {
		fmt.Println("Failed to read input.")
		return
	}

	if scanner.Err() != nil {
		fmt.Printf("Error: %v\n", scanner.Err())
		return
	}

	var wg sync.WaitGroup
	for serverName, database := range databaseConnections {
		wg.Add(1)
		go func(serverName string, database *sql.DB) {
			defer wg.Done()
			sendMessageToServer(serverName, database, message)
		}(serverName, database)
	}

	wg.Wait()
}

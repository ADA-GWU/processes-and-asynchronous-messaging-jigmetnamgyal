package main

import (
	"database/sql"
	"fmt"
	"log"
	"processes-and-asynchronous-messaging-jigmetnamgyal/db"
	"sync"
	"time"
)

var senderFullName = "ashfjkasn"

func readAvailableMessages(serverName string, db *sql.DB) {
	query := `
		SELECT ID, SENDER_NAME, MESSAGE, "CURRENT_TIME" FROM ASYNC_MESSAGE 
	  	WHERE RECEIVED_TIME IS NULL AND SENDER_NAME != $1 FOR UPDATE
	`

	rows, err := db.Query(query, senderFullName)

	if err != nil {
		log.Printf("Error reading messages from %s: %v", serverName, err)
		return
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal("Failed to close row")
		}
	}(rows)

	for rows.Next() {
		var id int
		var senderName string
		var message string
		var currentTime time.Time
		err = rows.Scan(&id, &senderName, &message, &currentTime)

		if err != nil {
			log.Printf("Error scanning rows from %s: %v", serverName, err)
			continue
		}

		fmt.Printf("Sender `%s` sent `%s` at time `%s`.\n", senderName, message, currentTime)

		updateQuery := "UPDATE ASYNC_MESSAGE SET RECEIVED_TIME = $1 WHERE ID = $2"

		_, err = db.Exec(updateQuery, time.Now(), id)

		if err != nil {
			log.Printf("Error updating RECEIVED_TIME for message from %s: %v", serverName, err)
		}
	}
}

func main() {
	databaseConnections := db.InitializeDatabaseConnections()

	for {
		var wg sync.WaitGroup

		for serverName, database := range databaseConnections {
			wg.Add(1)
			go func(serverName string, database *sql.DB) {
				defer wg.Done()
				readAvailableMessages(serverName, database)
			}(serverName, database)
		}

		wg.Wait()
	}

}

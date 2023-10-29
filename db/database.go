package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseInfo struct {
	ServerName       string
	ConnectionString string
}

func InitializeDatabaseConnections() map[string]*sql.DB {
	databaseInfo := []DatabaseInfo{
		{
			ServerName:       "Server1",
			ConnectionString: "postgres://assignment:assignment11@192.168.1.35:5432/assignment_db?sslmode=disable",
		},
	}

	databaseConnections := make(map[string]*sql.DB)

	for _, info := range databaseInfo {
		db := connectToDatabase(info.ServerName, info.ConnectionString)

		databaseConnections[info.ServerName] = db
	}

	return databaseConnections
}

func connectToDatabase(serverName string, connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error connecting to db")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to create db connection:" + err.Error())
	}

	return db
}

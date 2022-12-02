package main

import (
	"DCMS/api"
	"DCMS/db/sqlc"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver      = "postgres"
	serverAddress = "0.0.0.0:8080"
	dbSource      = "postgresql://root:secret@localhost:5432/DCMS?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot start the server... ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server... ", err)
	}
}

package main

import (
	"DCMS/api"
	db "DCMS/db/postgresql/sqlc"
	"DCMS/util"
	"DCMS/watcher"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go startServer(&wg)
	go watcher.StartWatching(&wg)
	wg.Wait()
}

func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot start the server... ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server... ", err)
	}
}

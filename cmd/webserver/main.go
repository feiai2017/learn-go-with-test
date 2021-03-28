package main

import (
	"github.com/Learn-go-with-tests/http_server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := http_server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := http_server.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

package main

import (
	"fmt"
	"github.com/Learn-go-with-tests/http_server"
	"log"
	"os"
)

const dbFileName = "../webserver/game.db.json"

func main() {
	store, close, err := http_server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	http_server.NewCli(store, os.Stdin).PlayPoker()
}

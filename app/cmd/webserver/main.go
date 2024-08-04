package main

import (
	"log"
	"net/http"

	poker "github.com/Abi-Liu/go-with-tests/app"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}

package main

import (
	"log"

	"github.com/nulfrost/ecom/cmd/api"
	"github.com/nulfrost/ecom/db"
)

func main() {

	// maybe put this in config somehow
	db, err := db.NewSQLiteStorage("./ecom.db")

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

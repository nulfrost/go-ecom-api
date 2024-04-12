package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/nulfrost/ecom/cmd/api"
	"github.com/nulfrost/ecom/config"
	"github.com/nulfrost/ecom/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		Net:                  "tcp",
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

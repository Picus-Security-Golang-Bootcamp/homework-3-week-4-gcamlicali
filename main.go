package main

import (
	postgres "DBHW/DB"
	//"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	//Reading ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	// Creating new db
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
		os.Exit(1)
	}
	log.Println("Postgres connected")

	postgres.ExecuteQueries(db)

}

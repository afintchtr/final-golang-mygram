package main

import (
	"final-golang-mygram/database"
	"final-golang-mygram/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }

	database.StartDB()
	r := routers.StartApp()
	r.Run(os.Getenv("PORT"))
}

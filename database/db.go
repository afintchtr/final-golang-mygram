package database

import (
	"final-golang-mygram/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	dbport := os.Getenv("DBPORT")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, dbport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	fmt.Println("Successfully connected to database")
	db.AutoMigrate(models.SocialMedia{}, models.Comment{}, models.Photo{}, models.User{})
}

func GetDB() *gorm.DB {
	return db
}

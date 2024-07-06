package config

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	// Connect to the "bank" database.
	godotenv.Load(".env")
	fmt.Println("Successfully loaded .env file")
	var fetchURL = os.Getenv("url")
	d, err := gorm.Open("postgres", fetchURL)
	if err != nil {
		log.Fatal(err)
	}
	db = d
	fmt.Print("Connected to the database\n")
}
func getDB() *gorm.DB {
	return db
}

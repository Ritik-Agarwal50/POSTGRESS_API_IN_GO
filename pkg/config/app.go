package config

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	fetchURL := "<PUT YOUR DATABSE URL HERE>"
	d, err := gorm.Open("postgres", fetchURL)

	if err != nil {
		log.Fatal(err)
	}
	db = d
	fmt.Print("Connected to the database\n")
}
func GetDB() *gorm.DB {
	return db
}

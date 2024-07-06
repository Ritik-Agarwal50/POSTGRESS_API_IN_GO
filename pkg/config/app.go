package config
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
var (
	db *sql.DB
)
func Connect() {
	// Connect to the "bank" database.
	godotenv.Load(".env")
	fmt.Println("Successfully loaded .env file")
	var fetchURL = os.Getenv("url")
	d, err := sql.Open("postgres", fetchURL)
	if err != nil {
		log.Fatal(err)
	}
	db = d
	fmt.Print("Connected to the database\n")
}
func getDB() *sql.DB {
	return db
}
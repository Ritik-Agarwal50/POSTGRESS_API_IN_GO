package main
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
func main() {
	// Connect to the "bank" database.
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Successfully loaded .env file")
	var fetchURL = os.Getenv("url")
	_, err := sql.Open("postgres", fetchURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Connected to the database\n")

	type movies struct{
		
	} 
}
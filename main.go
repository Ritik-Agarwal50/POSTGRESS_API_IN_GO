package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var url = "user=postgres.uukyjvzomuaxwnfjgatu password=0VwLOEtc4ZMZ65LM host=aws-0-us-east-1.pooler.supabase.com port=6543 dbname=postgres"
	_, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Connected to the database\n")
}

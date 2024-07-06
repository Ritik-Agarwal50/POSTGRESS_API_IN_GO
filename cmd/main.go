package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ritik-agarwal50/postgress_db/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, r))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

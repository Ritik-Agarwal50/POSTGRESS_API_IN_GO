package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ritik-agarwal50/postgress_db/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)
	http.Handle("/", r)
	port := 6543
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

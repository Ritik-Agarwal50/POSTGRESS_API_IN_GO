package routes

import (
	"github.com/gorilla/mux"
	"github.com/ritik-agarwal50/postgress_db/pkg/controllers"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{MovieId}", controllers.DeleteMovie).Methods("DELETE")

}

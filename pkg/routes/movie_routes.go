package routes

import (
	"github.com/gorilla/mux"
	"github.com/ritik-agarwal50/postgress_db/pkg/controllers"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.updateMovie).Methods("PUT")
	router.HandleFunc("/movie/{MovieId}", controllers.deleteMovie).Methods("DELETE")

}

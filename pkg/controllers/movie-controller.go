package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ritik-agarwal50/postgress_db/pkg/models"
	"github.com/ritik-agarwal50/postgress_db/pkg/utils"
)

var newMovie models.Movies

func GetMovie(w http.ResponseWriter, r *http.Request) {
	newMovies := models.GetAllMovies()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	newMovie, _ := models.GetMovieById(ID)
	res, _ := json.Marshal(newMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movies{}
	utils.ParseBody(r, CreateMovie)
	newMovie := CreateMovie.CreateMovie()
	res, _ := json.Marshal(newMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	deleteMovie := models.DeleteMovie(ID)
	res, _ := json.Marshal(deleteMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &models.Movies{}
	utils.ParseBody(r, updateMovie)
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	newMovie, db := models.GetMovieById(ID)
	if newMovie.Title != "" {
		newMovie.Title = updateMovie.Title
		newMovie.Write = updateMovie.Write
		newMovie.Publication = updateMovie.Publication
		db.Save(&newMovie)
		res, _ := json.Marshal(newMovie)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/ritik-agarwal50/postgress_db/pkg/models"

)

var newMovie models.Movie

func GetMovie(){}
func GetMovieById(){}
func CreateMovie(){}
func DeleteMovie(){}
func UpdateMovie(){}
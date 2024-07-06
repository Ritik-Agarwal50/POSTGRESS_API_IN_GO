package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ritik-agarwal50/postgress_db/pkg/config"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name        string `gorm:""json:name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func (b *Movie) CreateMovie() *Movie {
	db.NewRecord(b)
	db.Create(b)
	return b
}
func GetAllMovies() []Movie {
	var Movies []Movie
	db.Find(&Movies)
	return Movies
}
func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(ID int64) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}

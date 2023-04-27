package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var movies[] Movie



func main()  {
	r:=mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One",Director: &Director{Firstname: "John",Lastname: "Wick"}})
	movies = append(movies, Movie{ID: "2",Isbn: "45444", Title: "Movie Two",Director: &Director{Firstname: "Steve",Lastname: "Smith"}})

	//Create five different route and five different function
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting the Server at this Port\n")
	//logout in case of server dont listen
	log.Fatal(http.ListenAndServe(":8000",r))
	
	
}
func getMovies(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}

	}

}

func createMovie(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type","json/application")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r* http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "json/application")
	//params
	params := mux.Vars(r)

	//range over movies
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
	//delete the movie with id that you have sent
	//add a new movie -  the movie that we send in the body of postman



}

func deleteMovie(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}



// Movies Details
type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

// Director details
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

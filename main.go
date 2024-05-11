package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// construct the type

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Movie 1", Director: "Liton"})
	r.HandleFunc("/movies", getMovies).Methods("GET")

	fmt.Println("Starting server at port 80000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

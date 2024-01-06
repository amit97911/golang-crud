package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    "json:id"
	Isbn     string    "json:isbn"
	Title    string    "json:title"
	Director *Director "json:director"
}

type Director struct {
	FirstName string "json:firstname"
	LastName  string "json:lastname"
}

var (
	movies []Movie
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "isbn1", Title: "title1", Director: &Director{FirstName: "firstName1", LastName: "lastName1"}})
	movies = append(movies, Movie{ID: "2", Isbn: "isbn", Title: "title2", Director: &Director{FirstName: "firstName2", LastName: "lastName2"}})
	movies = append(movies, Movie{ID: "3", Isbn: "isbn", Title: "title3", Director: &Director{FirstName: "firstName3", LastName: "lastName3"}})
	movies = append(movies, Movie{ID: "4", Isbn: "isbn", Title: "title4", Director: &Director{FirstName: "firstName4", LastName: "lastName4"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	// router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// router.HandleFunc("/movies", createMovie).Methods("POST")
	// router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started and running on ", srv.Addr)
	defer log.Fatal(srv.ListenAndServe())
}

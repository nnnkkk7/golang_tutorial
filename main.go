package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Tittle string  `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewDecoder(w).Encode(item)
			return
		}
		json.NewDecoder(w).Encode(&Book{})
	}

}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewDecoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) 
			var book Book
	        _ = json.NewDecoder(r.Body).Decode(&book)
	        book.ID = params["id"]
			books = append(books, book)
			json.NewDecoder(w).Encode(book)
			return
		}	
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) 
			break
		}
		json.NewDecoder(w).Encode(books)
}

func main() {

	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "4242", Tittle: "book one", Author: &Author{Firstname: "john", Lastname: "noha"}})

	r.HandleFunc("api/books", getBooks).Methods("GET")
	r.HandleFunc("api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("api/books", createBook).Methods("POST")
	r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

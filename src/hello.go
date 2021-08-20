package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

//BOOK Struct (Model)
type Book struct {
	ID string `json: "id"`
	Isbn  string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as slice Book struct
var books []Book


// Get all books
func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
// Get single book
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
//	loop over the books and find one
	for _, item := range books{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
// Create a new book`
func createBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = uuid.New().String()
	books = append(books, book)
	json.NewEncoder(w).Encode(&book)
}
// update a book
func updateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(&book)
			return
		}
	}
	json.NewEncoder(w).Encode(&books)
}
// delete book
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(&books)
}

func main()  {


	//Mock data
	books = append(books, Book{ID:"1", Isbn: "12312", Title: "Book 1", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID:"2", Isbn: "433432", Title: "Book 2", Author: &Author{Firstname: "Mike", Lastname: "Smith"}})


	//Init router
	// :=  set type by the value
	r := mux.NewRouter()
	//route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
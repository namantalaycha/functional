package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"math/rand"
	"net/http"
	"strconv"
)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Books []Book



func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
    fmt.Println("hello")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params :=  chi.URLParam(r, "id")

	for _, item := range Books {
		if item.ID == params {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Fprintf(w,"No Match Found!")
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
	fmt.Fprintf(w,"book is created")

}

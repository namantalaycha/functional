package app

import (
	"fmt"
	"functional/middleware"
	"functional/service"
	"github.com/go-chi/chi"
	"net/http"
)


func Startapp( ){
	r := chi.NewRouter()
	service.Books = append(service.Books, service.Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &service.Author{Firstname: "John", Lastname: "Doe"}})
	service.Books = append(service.Books, service.Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &service.Author{Firstname: "Steve", Lastname: "Smith"}})
	r.Use(middleware.Auth)
	r.Get("/books", service.GetBooks)
	r.Get("/books/{id}", service.GetBook)
	r.Post("/books",service.CreateBook)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}

}

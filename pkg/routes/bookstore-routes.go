package routes

import (
	"github.com/gorilla/mux"
	"github.com/shrey/go_bookstore/pkg/controllers"
)

// github.com/shrey/go_bookstore

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("book/", controllers.CreateBook).Methods("POST")

	router.HandleFunc("book/", controllers.GetBook).Methods("GET")

	router.HandleFunc("book/{bookId}", controllers.GetBookById).Methods("GET")

	router.HandleFunc("book/{bookId}", controllers.UpdateBook).Methods("PUT")

	router.HandleFunc("book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}

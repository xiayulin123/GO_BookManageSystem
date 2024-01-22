package routes

import (
	"github.com/gorilla/mux"
	"github.com/xiayulin123/Go-bookstore/pkg/controllers"
)

var BookStroeRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")

}

package routes

import (
	"github.com/gorilla/mux"
	"mongo-api/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookbyTitle).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}

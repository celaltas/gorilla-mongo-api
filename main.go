package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mongo-api/routes"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	if err := http.ListenAndServe(":9000", r); err != nil {
		fmt.Println(err)
	}
}

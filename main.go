package main

import (
	"log"
	"net/http"

	"github.com/danigilang17/ptzmi/database"
	"github.com/danigilang17/ptzmi/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/products", routes.CreateProduct).Methods("POST")
	r.HandleFunc("/products", routes.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", routes.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", routes.DeleteProduct).Methods("DELETE")

	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

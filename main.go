package main

import (
	"log"
	"net/http"

	"github.com/danigilang17/tokoonline/database"
	"github.com/danigilang17/tokoonline/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	r := mux.NewRouter()
	// Product Routes
	r.HandleFunc("/products", routes.CreateProduct).Methods("POST")
	r.HandleFunc("/products", routes.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", routes.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", routes.DeleteProduct).Methods("DELETE")

	// Order Routes
	r.HandleFunc("/orders", routes.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", routes.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", routes.UpdateOrderStatus).Methods("PUT")

	// Authentication Routes
	r.HandleFunc("/register", routes.Register).Methods("POST")
	r.HandleFunc("/login", routes.Login).Methods("POST")
	r.HandleFunc("/profile", routes.Profile).Methods("GET")
	r.HandleFunc("/reset-password", routes.ResetPassword).Methods("POST")
	r.HandleFunc("/logout", routes.Logout).Methods("POST")

	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

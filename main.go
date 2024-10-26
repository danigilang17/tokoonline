package main

import (
	"log"
	"net/http"

	"github.com/danigilang17/tokoonline/database"
	"github.com/danigilang17/tokoonline/middleware"
	"github.com/danigilang17/tokoonline/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	r := mux.NewRouter()

	// Routes tanpa autentikasi
	r.HandleFunc("/register", routes.Register).Methods("POST")
	r.HandleFunc("/login", routes.Login).Methods("POST")
	r.HandleFunc("/reset-password", routes.ResetPassword).Methods("POST")
	r.HandleFunc("/logout", routes.Logout).Methods("POST")

	// Routes yang memerlukan autentikasi JWT
	authRoutes := r.PathPrefix("/").Subrouter()
	authRoutes.Use(middleware.JWTAuth)
	authRoutes.HandleFunc("/profile", routes.Profile).Methods("GET")
	authRoutes.HandleFunc("/products", routes.CreateProduct).Methods("POST")
	authRoutes.HandleFunc("/products", routes.GetProducts).Methods("GET")
	authRoutes.HandleFunc("/products/{id}", routes.UpdateProduct).Methods("PUT")
	authRoutes.HandleFunc("/products/{id}", routes.DeleteProduct).Methods("DELETE")
	authRoutes.HandleFunc("/orders", routes.CreateOrder).Methods("POST")
	authRoutes.HandleFunc("/orders", routes.GetOrders).Methods("GET")
	authRoutes.HandleFunc("/orders/{id}", routes.UpdateOrderStatus).Methods("PUT")

	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

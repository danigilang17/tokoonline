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

	// Public routes
	r.HandleFunc("/register", routes.Register).Methods("POST")
	r.HandleFunc("/login", routes.Login).Methods("POST")

	// Protected routes
	r.HandleFunc("/products", middleware.JWTAuth(routes.CreateProduct)).Methods("POST")
	r.HandleFunc("/products", middleware.JWTAuth(routes.GetProducts)).Methods("GET")
	r.HandleFunc("/products/{id}", middleware.JWTAuth(routes.UpdateProduct)).Methods("PUT")
	r.HandleFunc("/products/{id}", middleware.JWTAuth(routes.DeleteProduct)).Methods("DELETE")

	// Order Routes
	r.HandleFunc("/orders", middleware.JWTAuth(routes.CreateOrder)).Methods("POST")
	r.HandleFunc("/orders", middleware.JWTAuth(routes.GetOrders)).Methods("GET")
	r.HandleFunc("/orders/{id}", middleware.JWTAuth(routes.UpdateOrderStatus)).Methods("PUT")

	// Other protected routes
	r.HandleFunc("/profile", middleware.JWTAuth(routes.Profile)).Methods("GET")
	r.HandleFunc("/reset-password", middleware.JWTAuth(routes.ResetPassword)).Methods("POST")
	r.HandleFunc("/logout", middleware.JWTAuth(routes.Logout)).Methods("POST")

	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

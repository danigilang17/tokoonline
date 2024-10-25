package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danigilang17/tokoonline/database"
	"github.com/danigilang17/tokoonline/models"
	"github.com/gorilla/mux"
)

// CreateOrder creates a new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	result, err := database.DB.Exec("INSERT INTO orders (product_id) VALUES (?)", order.ProductID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order.ID = int(id) // Ubah ini, asumsikan ID di struct Order bertipe int64
	order.Status = "pending"
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrders retrieves all orders
func GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.ProductID, &order.Status, &order.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}
	json.NewEncoder(w).Encode(orders)
}

// UpdateOrderStatus updates the status of an order
func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert idStr to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	_, err = database.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", order.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order.ID = int(id)
	json.NewEncoder(w).Encode(order)
}

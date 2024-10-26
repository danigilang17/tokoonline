package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danigilang17/tokoonline/database"
	"github.com/danigilang17/tokoonline/models"
	"github.com/gorilla/mux"
)

// CreateOrder handles the creation of a new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	// Validate User ID
	var userExists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", order.UserID).Scan(&userExists)
	if err != nil || !userExists {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Retrieve product price
	var price float64
	err = database.DB.QueryRow("SELECT price FROM products WHERE id = ?", order.ProductID).Scan(&price)
	if err != nil {
		http.Error(w, "Invalid product_id", http.StatusBadRequest)
		return
	}

	// Calculate total price
	total := price * float64(order.Qty)

	// Insert order with calculated total
	result, err := database.DB.Exec("INSERT INTO orders (product_id, user_id, qty, status) VALUES (?, ?, ?, ?)", order.ProductID, order.UserID, order.Qty, "pending")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set total price in response
	order.Total = total
	order.Status = "pending"

	id, _ := result.LastInsertId()
	order.ID = int(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT orders.id, orders.product_id, orders.user_id, orders.qty, orders.status, products.price 
		FROM orders 
		INNER JOIN products ON orders.product_id = products.id
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var price float64

		if err := rows.Scan(&order.ID, &order.ProductID, &order.UserID, &order.Qty, &order.Status, &price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Calculate total price for each order
		order.Total = price * float64(order.Qty)
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

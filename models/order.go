package models

type Order struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	UserID    int     `json:"user_id"`
	Qty       int     `json:"qty"`
	Status    string  `json:"status"`
	Total     float64 `json:"total,omitempty"` // Field untuk total harga
}

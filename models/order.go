package models

type Order struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Status    string `json:"status"` // 'pending', 'processed', 'completed'
	CreatedAt string `json:"created_at"`
}

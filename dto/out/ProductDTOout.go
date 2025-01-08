package out

import "time"

type ProductResponse struct {
	ID           string    `json:"id"`
	ProductName  string    `json:"product_name"`
	Price        float32   `json:"price`
	ProductStock int64     `json:"product_stock"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

package in

type ProductRequest struct {
	ID    string   `json:"id"`
	Name  string  `json:"product_name"`
	Price float32 `json:"price"`
	Stock int64   `json:"stock"`
}

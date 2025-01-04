package model

import "time"

type Product struct {
	ID        string    `bson:"_id" json:"id"`
	Name      string    `bson:"product_name" json:"product_name`
	Price     float32   `bson :"price" json:"price"`
	Stock     int64     `bsom:"product_stock json:"product_stock"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Deleted   bool      `bson:"deleted" json:"deleted"`
}

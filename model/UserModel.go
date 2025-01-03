package model

import (
	"time"
)

type User struct {
	ID        int64    `bson:"_id,omitempty" json:"id"`
	Username  string    `bson:"username" json:"username"`
	Password  string    `bson:"password" json:"password"`
	FirstName string    `bson:"first_name" json:"firstName"`
	LastName  string    `bson:"last_name" json:"lastName"`
	Gender    string    `bson:"gender" json:"gender"`
	Phone     string    `bson:"phone" json:"phone"`
	Email     string    `bson:"email" json:"email"`
	Address   string    `bson:"address" json:"address"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt"`
}

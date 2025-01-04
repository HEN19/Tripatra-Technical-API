package dao

import (
	"context"
	"errors"
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type productDAO struct {
	AbstractDAO
}

var ProductDAO = productDAO{}.New()

func (input productDAO) New() (output productDAO) {
	output.TableName = "products"
	output.FileName = "ProductDAO.go"
	return
}

func (p productDAO) InsertProduct(inputStruct model.Product) (*mongo.InsertOneResult, error) {
	collection := config.GetMongoCollection("mydatabase", "products")

	product := bson.M{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := collection.InsertOne(ctx, product))
	if err != nil {
		return nil, err
	}

	return id, nil
}
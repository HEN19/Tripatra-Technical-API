package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	collection := config.GetMongoCollection("tripatra", "products")

	product := bson.M{
		"product_name":  inputStruct.Name,
		"price":         inputStruct.Price,
		"product_stock": inputStruct.Stock,
		"created_at":    time.Now(),
		"updated_at":    time.Now(),
		"deleted":       false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (p productDAO) GetListProduct() ([]model.Product, error) {
	collection := config.GetMongoCollection("tripatra", "products")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"deleted": false}

	var products []model.Product
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return products, nil
}

func (p productDAO) GetDetailProduct(id string) (model.Product, error) {
	collection := config.GetMongoCollection("tripatra", "products")

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"id": objectID, "deleted": false}

	fmt.Println(filter)
	var product model.Product
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (p productDAO) UpdateProduct(inputStruct model.Product) (*mongo.UpdateResult, error) {
	collection := config.GetMongoCollection("tripatra", "products")

	objectID, _ := primitive.ObjectIDFromHex(inputStruct.ID)
	filter := bson.M{"id": objectID}

	product := bson.M{
		"$set": bson.M{
			"product_name":  inputStruct.Name,
			"price":         inputStruct.Price,
			"product_stock": inputStruct.Stock,
			"updated_at":    time.Now(),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p productDAO) DeleteProduct(id string) (*mongo.UpdateResult, error) {
	collection := config.GetMongoCollection("tripatra", "products")

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"id": objectID}

	product := bson.M{
		"$set": bson.M{
			"deleted":    true,
			"updated_at": time.Now(),
		}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

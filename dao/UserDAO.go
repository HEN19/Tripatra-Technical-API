package dao

import (
	"context"
	"errors"
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDAO struct {
	AbstractDAO
}

var UserDAO = userDAO{}.New()

func (input userDAO) New() (output userDAO) {
	output.TableName = "users"
	output.FileName = "UserDAO.go"
	return
}

func (u userDAO) InsertUser(inputStruct model.User) (*mongo.InsertOneResult, error) {
	collection := config.GetMongoCollection("tripatra", "users")

	user := bson.M{
		"username":   inputStruct.Username,
		"password":   inputStruct.Password,
		"first_name": inputStruct.FirstName,
		"last_name":  inputStruct.LastName,
		"gender":     inputStruct.Gender,
		"phone":      inputStruct.Phone,
		"email":      inputStruct.Email,
		"address":    inputStruct.Address,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted":    false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (u userDAO) LoginCheck(user model.User) (model.User, error) {
	collection := config.GetMongoCollection("tripatra", "users")

	// Filter for matching username and password
	filter := bson.M{
		"username": user.Username,
		"password": user.Password,
	}

	var result model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, errors.New("invalid username or password")
		}
		return model.User{}, err
	}

	return result, nil
}

func (u userDAO) GetUserProfile(id string) (model.User, error) {
	collection := config.GetMongoCollection("tripatra", "users")

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}

	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

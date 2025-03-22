package mongodb

import (
	"context"
	"go_server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
	userCollection = client.Database("lowcode").Collection("users")
}

func CreateUser(user *model.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
} 
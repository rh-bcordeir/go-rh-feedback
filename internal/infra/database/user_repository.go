package database

import (
	"context"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserDB struct {
	collection *mongo.Collection
}

func NewUserDB(client *mongo.Client) *UserDB {
	return &UserDB{
		collection: client.Database("feedback_db").Collection("users"),
	}
}

func (userDB *UserDB) FindByEmail(email string) (*entity.User, error) {

	var userResult entity.User
	filter := bson.D{{"email", email}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := userDB.collection.FindOne(ctx, filter).Decode(&userResult)

	if err != nil {
		return nil, err
	}

	return &userResult, nil
}

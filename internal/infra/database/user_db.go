package database

import (
	"context"
	"errors"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	filter := bson.D{{Key: "email", Value: email}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := userDB.collection.FindOne(ctx, filter).Decode(&userResult)

	if err != nil {
		return nil, errors.New("no users were found")
	}

	return &userResult, nil
}

func (userDB *UserDB) SaveUser(user *entity.User) error {
	_, err := userDB.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already taken")
	}
	user.ID = primitive.NewObjectID()
	_, err = userDB.collection.InsertOne(context.Background(), user)
	return err
}

package database

import "go.mongodb.org/mongo-driver/v2/mongo"

type FeedbackRepository struct {
	client *mongo.Client
}

func NewFeedbackRepository(client *mongo.Client) *FeedbackRepository {
	return &FeedbackRepository{
		client: client,
	}
}

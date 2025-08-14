package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Stage struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" `
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}

func NewStage(title, description string) *Stage {
	return &Stage{
		Title:       title,
		Description: description,
	}
}

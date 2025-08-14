package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Position struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
}

func NewPosition(title string) *Position {
	return &Position{
		Title: title,
	}
}

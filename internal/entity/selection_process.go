package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SelectionProcess struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Stages    []string           `bson:"stages" json:"stages"`
	Feedbacks []Feedback         `bson:"feedbacks" json:"feedbacks"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

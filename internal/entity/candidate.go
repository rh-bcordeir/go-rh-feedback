package entity

import (
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
	Position  primitive.ObjectID `bson:"position"`
	Feedbacks []Feedback         `bson:"feedbacks"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func NewCandidate(name, email, phone string, position string) (*Candidate, error) {
	positionId, err := primitive.ObjectIDFromHex(position)

	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid position ID")
	}

	return &Candidate{
		Name:      name,
		Email:     email,
		Phone:     phone,
		Position:  positionId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

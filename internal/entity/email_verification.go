package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailVerification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	User      primitive.ObjectID `bson:"user_id," json:"user_id"`
	Email     string             `bson:"email" json:"email"`
	Token     string             `bson:"token"`
	ExpiresAt time.Time          `bson:"expires_at"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

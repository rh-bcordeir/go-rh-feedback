package entity

import (
	"time"

	"github.com/google/uuid"
)

type EmailVerification struct {
	ID        uuid.UUID
	User      uuid.UUID
	Email     string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

package entity

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Open     Status = "open"
	Ongoing  Status = "ongoing"
	Closed   Status = "closed"
	Canceled Status = "canceled"
)

type HiringProcess struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CandidateID uuid.UUID `gorm:"not null" `
	PositionID  uint      `gorm:"not null" `
	Candidate   Candidate `gorm:"foreignKey:CandidateID"`
	Position    Position  `gorm:"foreignKey:PositionID"`
	Status      Status    `gorm:"type:text;default:'open'" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime;<-:create"`
	UpdatedAt   time.Time
}

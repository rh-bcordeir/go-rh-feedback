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
	CandidateID uuid.UUID `gorm:"not null" json:"candidate_id"`
	PositionID  uint      `gorm:"not null" json:"position_id"`
	Candidate   Candidate `gorm:"foreignKey:CandidateID" json:"-"`
	Position    Position  `gorm:"foreignKey:PositionID" json:"position"`
	Status      Status    `gorm:"type:text;default:'open'" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

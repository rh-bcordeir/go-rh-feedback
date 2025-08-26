package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CandidatePosition struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" `
	CandidateID uuid.UUID `gorm:"not null" `
	PositionID  uint      `gorm:"not null" `
	Candidate   Candidate `gorm:"foreignKey:CandidateID"`
	Position    Position  `gorm:"foreignKey:PositionID"`
}

func (c *CandidatePosition) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

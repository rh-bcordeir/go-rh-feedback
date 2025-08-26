package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Candidate struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string
	Email     string
	Phone     string
	Positions []Position `gorm:"many2many:candidate_positions;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Candidate) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

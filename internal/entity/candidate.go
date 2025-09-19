package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Candidate struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string     `json:"name" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null;uniqueIndex"`
	Phone     string     `json:"phone"`
	Positions []Position `gorm:"many2many:hiring_processes;" json:"positions,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (c *Candidate) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

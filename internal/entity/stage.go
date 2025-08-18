package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stage struct {
	ID          uuid.UUID
	Title       string
	Description string
}

func (f *Stage) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return nil
}

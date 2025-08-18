package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Position struct {
	ID    uuid.UUID
	Title string
}

func (f *Position) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return nil
}

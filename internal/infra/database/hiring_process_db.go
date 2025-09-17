package database

import (
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"gorm.io/gorm"
)

type HiringProcessDB struct {
	db *gorm.DB
}

func NewHiringProcessDB(db *gorm.DB) *HiringProcessDB {
	return &HiringProcessDB{
		db: db,
	}
}

func (h *HiringProcessDB) SaveHiringProcess(hiringProcess *entity.HiringProcess) error {
	return h.db.Create(hiringProcess).Error
}

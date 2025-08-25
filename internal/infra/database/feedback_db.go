package database

import (
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"gorm.io/gorm"
)

type FeedbackDB struct {
	db *gorm.DB
}

func NewFeedbackDB(db *gorm.DB) *FeedbackDB {
	return &FeedbackDB{db: db}
}

func (f *FeedbackDB) SaveFeedback(feedback *entity.Feedback) error {
	return f.db.Create(feedback).Error
}

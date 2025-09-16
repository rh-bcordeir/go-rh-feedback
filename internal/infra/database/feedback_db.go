package database

import (
	"errors"

	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"github.com/google/uuid"
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

func (f *FeedbackDB) GetAllFeedbacks() ([]entity.Feedback, error) {
	var feedbacks []entity.Feedback
	if err := f.db.Model(&entity.Feedback{}).Preload("Interviewer").Preload("Candidate").Preload("Stage").Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (f *FeedbackDB) DeleteFeedback(feedbackId string) error {
	feedbackUUID, err := uuid.Parse(feedbackId)
	if err != nil {
		return errors.New("invalid feedback ID")
	}

	return f.db.Delete(&entity.Feedback{ID: feedbackUUID}).Error
}

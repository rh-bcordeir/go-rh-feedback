package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	InterviewerID uuid.UUID `gorm:"type:uuid" json:"interviewer_id"`
	Interviewer   User      `gorm:"foreignKey:InterviewerID" json:"-"`
	CandidateID   uuid.UUID `gorm:"type:uuid" json:"candidate_id"`
	Candidate     Candidate `gorm:"foreignKey:CandidateID" json:"-"`
	StageID       uint      `gorm:"not null"`
	Stage         Stage     `gorm:"foreignKey:StageID" json:"-"`
	Comments      string    `gorm:"type:text" json:"comments"`
	Score         int       `json:"score"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (f *Feedback) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return nil
}

func NewFeedback(interviewerID, candidateID string, stageID uint, comments string, score int) *Feedback {
	interviewerUUID := uuid.MustParse(interviewerID)
	candidateUUID := uuid.MustParse(candidateID)

	return &Feedback{
		InterviewerID: interviewerUUID,
		CandidateID:   candidateUUID,
		StageID:       stageID,
		Comments:      comments,
		Score:         score,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

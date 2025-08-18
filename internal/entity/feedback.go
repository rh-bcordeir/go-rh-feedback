package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	InterviewerID uuid.UUID `gorm:"type:uuid" json:"interviewer_id"`
	Interviewer   User      `gorm:"foreignKey:InterviewerID"`
	CandidateID   uuid.UUID `gorm:"type:uuid" json:"candidate_id"`
	Candidate     Candidate `gorm:"foreignKey:CandidateID"`
	StageID       uuid.UUID `gorm:"type:uuid"`
	Stage         Stage     `gorm:"foreignKey:StageID"`
	Comments      string    `json:"comments"`
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

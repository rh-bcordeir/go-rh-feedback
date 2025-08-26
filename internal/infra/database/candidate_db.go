package database

import (
	"errors"

	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"gorm.io/gorm"
)

type CandidateDB struct {
	db *gorm.DB
}

func NewCandidateDB(db *gorm.DB) *CandidateDB {
	return &CandidateDB{db: db}
}

func (c *CandidateDB) CreateCandidate(candidate *entity.Candidate, positionId uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(&entity.Candidate{}).
			Where("email = ?", candidate.Email).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("candidate with this email already exists")
		}

		// 2. se positionId foi informado, valida antes de criar o candidate
		if positionId != 0 {
			var position entity.Position
			if err := tx.First(&position, "id = ?", positionId).Error; err != nil {
				return errors.New("position not found")
			}

			if err := tx.Create(candidate).Error; err != nil {
				return err
			}

			if err := tx.Create(&entity.CandidatePosition{
				CandidateID: candidate.ID,
				PositionID:  position.ID,
			}).Error; err != nil {
				return err
			}

			return nil
		}

		return tx.Create(candidate).Error
	})
}

func (c *CandidateDB) GetAllCandidates() ([]entity.Candidate, error) {
	var candidates []entity.Candidate
	if err := c.db.Model(&entity.Candidate{}).Preload("Positions").Find(&candidates).Error; err != nil {
		return nil, err
	}
	return candidates, nil
}

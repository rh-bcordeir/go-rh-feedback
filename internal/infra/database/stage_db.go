package database

import (
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"gorm.io/gorm"
)

type StageDB struct {
	db *gorm.DB
}

func NewStageDB(db *gorm.DB) *StageDB {
	return &StageDB{db: db}
}

func (s *StageDB) GetAllStages() ([]entity.Stage, error) {
	var stages []entity.Stage
	if err := s.db.Model(&entity.Stage{}).Distinct().Find(&stages).Error; err != nil {
		return nil, err
	}

	return stages, nil
}

func (s *StageDB) CreateStage(stage *entity.Stage) error {
	return s.db.Create(stage).Error
}

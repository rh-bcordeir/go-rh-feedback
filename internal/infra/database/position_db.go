package database

import (
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"gorm.io/gorm"
)

type PositionDB struct {
	db *gorm.DB
}

func NewPositionDB(db *gorm.DB) *PositionDB {
	return &PositionDB{db: db}
}

func (p *PositionDB) GetAllPositions() ([]entity.Position, error) {
	var positions []entity.Position
	if err := p.db.Model(&entity.Position{}).Distinct().Find(&positions).Error; err != nil {
		return nil, err
	}

	return positions, nil
}

func (p *PositionDB) CreatePosition(position *entity.Position) error {
	return p.db.Create(position).Error
}

func (p *PositionDB) FindByID(id uint) (*entity.Position, error) {
	var position entity.Position
	err := p.db.First(&position, "id = ?", id).Error
	return &position, err
}

func (p *PositionDB) UpdatePosition(position *entity.Position) error {
	_, err := p.FindByID(position.ID)
	if err != nil {
		return err
	}
	return p.db.Save(position).Error
}

func (p *PositionDB) DeletePosition(id uint) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.db.Delete(&entity.Position{}, id).Error
}

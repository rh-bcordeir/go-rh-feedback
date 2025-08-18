package database

import (
	"errors"

	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) SaveUser(user *entity.User) error {

	var count int64
	if err := u.db.Model(&entity.User{}).Where("email = ?", user.Email).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("email already taken")
	}
	return u.db.Create(user).Error
}

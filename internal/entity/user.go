package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	INTERVIEWER Role = "interviewer"
	RECRUITER   Role = "recruiter"
	ADMIN       Role = "admin"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name          string    `json:"name"`
	Email         string    `gorm:"uniqueIndex;size:255" json:"email"`
	EmailVerified bool      `gorm:"default:false" json:"email_verified"`
	Password      string    `json:"-"`
	Role          Role      `gorm:"type:text;default:'interviewer'" json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ValidateEmail(email string) error {
	if !strings.HasSuffix(strings.ToLower(email), "@redhat.com") {
		return errors.New("email must be a redhat.com email")
	}
	return nil
}

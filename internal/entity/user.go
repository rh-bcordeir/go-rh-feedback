package entity

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	INTERVIEWER Role = "Interviewer"
	MANAGER     Role = "Manager"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name" `
	Email         string             `bson:"email" `
	EmailVerified time.Time          `bson:"emailVerified" `
	Password      string             `bson:"password" `
	Role          Role               `bson:"role" `
	CreatedAt     time.Time          `bson:"createdAt" `
	UpdatedAt     time.Time          `bson:"updatedAt" `
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:      name,
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	fmt.Println(u.Password)
	fmt.Println(password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ValidateEmail(email string) error {
	if !strings.HasSuffix(strings.ToLower(email), "@redhat.com") {
		return errors.New("email must be a redhat.com email")
	}

	return nil
}

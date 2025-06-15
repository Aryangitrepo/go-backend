package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int `gorm:"primaryKey;autoIncrement:true"`
	Name     string
	Role     string
	Email    string `gorm:"unique;not null"`
	Password string
	Pic      string
}
type UserFunc interface {
	AddUser(user *User) error
	FindUser(credentials LoginJson) (*User, error)
}

// User operations
func (u *User) AddUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	return nil
}

func (u *User) FindUser(credentials LoginJson) (*User, error) {
	var user User
	err := db.Where("email = ? AND password = ?", credentials.Email, credentials.Password).
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with provided credentials")
		}
		return nil, fmt.Errorf("finding user: %w", err)
	}
	return &user, nil
}

package user

import (
	"errors"
	"strings"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}

func NewUser(name, email string) (*User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name is required")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}

	return &User{
		Name:  name,
		Email: email,
	}, nil
}

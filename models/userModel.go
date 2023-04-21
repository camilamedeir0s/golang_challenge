package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	FirstName string
	LastName string
	Email string
	Password string
	Phone string
	UserStatus string
  }
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"unique;not null" validate:"required,min=3,max=191,my-custom-valid" json:"name" json:"name"`
	Email  string `gorm:"unique;not null" validate:"required,email" json:"email" json:"email"`
	Status string `json:"status"`

	Wallet []Wallet
}

package models

import (
	"card-game/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string            `gorm:"unique;not null" validate:"required,min=3,max=191,user-exist" json:"name"`
	Email  string            `gorm:"unique;not null" validate:"required,email,user-exist" json:"email"`
	Status consts.UserStatus `gorm:"not null" validate:"required,oneof=1 2" json:"status"`

	Wallet []Wallet
}

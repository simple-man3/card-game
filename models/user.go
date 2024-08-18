package models

import (
	"card-game/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string            `gorm:"unique;not null" json:"name"`
	Email  string            `gorm:"unique;not null" json:"email"`
	Status consts.UserStatus `gorm:"not null" json:"status"`

	Wallet []Wallet
}

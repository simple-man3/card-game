package models

import (
	"card-game/consts"
	"time"
)

type User struct {
	ID     uint              `gorm:"primarykey" json:"id"`
	Name   string            `gorm:"unique;not null" json:"name"`
	Email  string            `gorm:"unique;not null" json:"email"`
	Status consts.UserStatus `gorm:"not null" json:"status"`

	Wallet []Wallet `json:"wallets"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package models

import (
	"card-game/consts"
	"time"
)

var AuthUser *User

type User struct {
	ID       uint              `gorm:"primarykey" json:"id"`
	Name     string            `gorm:"unique;not null" json:"name"`
	Email    string            `gorm:"unique;not null" json:"email"`
	Password string            `gorm:"not null" json:"-"`
	Status   consts.UserStatus `gorm:"not null" json:"status"`

	Wallet *Wallet `json:"wallet,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

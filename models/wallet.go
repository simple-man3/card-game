package models

import (
	"time"
)

type Wallet struct {
	ID      uint    `gorm:"primarykey" json:"id"`
	UserId  uint    `gorm:"index" json:"user_id"`
	Balance float64 `gorm:"type:decimal(10,2);" json:"balance"`

	User         User          `json:"user"`
	Transactions []Transaction `json:"transactions"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

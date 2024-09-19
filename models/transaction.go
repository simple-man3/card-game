package models

import (
	"card-game/consts"
	"time"
)

type Transaction struct {
	ID       uint `gorm:"primarykey"`
	WalletId uint `gorm:"index"`
	Sum      float64
	Action   consts.TrxAction

	Wallet Wallet `json:"wallet"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	WalletId uint `gorm:"index"`
	Sum      float64
	Action   string

	Wallet Wallet
}

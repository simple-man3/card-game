package models

import "time"

type Transaction struct {
	ID       uint `gorm:"primarykey"`
	WalletId uint `gorm:"index"`
	Sum      float64
	Action   string

	Wallet Wallet `json:"wallet"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

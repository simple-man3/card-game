package models

import "time"

type Wallet struct {
	ID      uint `gorm:"primarykey"`
	UserId  uint `gorm:"index"`
	Balance float64

	User         User
	Transactions []Transaction

	CreatedAt time.Time
	UpdatedAt time.Time
}

package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserId  uint `gorm:"index"`
	Balance float64

	User         User
	Transactions []Transaction
}

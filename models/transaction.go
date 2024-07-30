package models

import (
	"card-game/database"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	WalletId uint `gorm:"index"`
	Sum      float64
	Action   string

	Wallet Wallet
}

func CreateTransaction(trx *Transaction) (*Transaction, error) {
	db := database.DBConn

	res := db.Create(&trx)
	if res.Error != nil {
		return nil, res.Error
	}

	return trx, nil
}

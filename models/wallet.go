package models

import (
	"card-game/database"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserId  uint `gorm:"index"`
	Balance float64

	User         User
	Transactions []Transaction
}

func CreateWallet(wallet *Wallet) (*Wallet, error) {
	db := database.DBConn

	res := db.Create(wallet)

	if res.Error != nil {
		return nil, res.Error
	}

	return wallet, nil
}

func PutMoney(amount float64, wallet *Wallet) error {
	db := database.DBConn

	wallet.Balance += amount

	res := db.Save(&wallet)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func TakeMoney(amount float64, wallet *Wallet) error {
	db := database.DBConn

	wallet.Balance -= amount

	res := db.Save(&wallet)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

package services

import (
	"card-game/database"
	"card-game/models"
)

func CreateTransaction(trx *models.Transaction) (*models.Transaction, error) {
	db := database.DBConn

	res := db.Create(&trx)
	if res.Error != nil {
		return nil, res.Error
	}

	return trx, nil
}

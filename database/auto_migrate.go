package database

import (
	"card-game/models"
	"errors"
)

func AutoMigrate() error {
	if DBConn == nil {
		return errors.New("database connection is nil")
	}

	if err := DBConn.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Transaction{},
	); err != nil {
		return err
	}

	return nil
}

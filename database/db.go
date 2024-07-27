package database

import (
	"card-game/config"
	"card-game/models"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	DBConn  *gorm.DB
	once    sync.Once
	connErr error
)

func Connect() error {
	env, _ := config.GetInstanceEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DbUsername, env.DbPassword, env.DbHost, env.DbPort, env.DbDatabase)

	once.Do(func() {
		DBConn, connErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})

	if connErr != nil {
		return connErr
	}

	return nil
}

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

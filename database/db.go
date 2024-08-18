package database

import (
	"card-game/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() error {
	env, _ := config.GetInstanceEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", env.DbUsername, env.DbPassword, env.DbHost, env.DbPort, env.DbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DBConn = db

	return nil
}

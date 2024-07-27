package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Email  string `gorm:"unique"`
	Status string

	Wallet []Wallet
}

package config

import (
	"github.com/aliirsyaadn/neobank-technical-test/model"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.UserOTP{},
		&model.Transaction{},
		&model.TransferRecord{},
	)
	if err != nil {
		panic(err)
	}
}

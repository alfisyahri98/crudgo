package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projectgo/model"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}

	return db
}

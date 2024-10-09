package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDatabaseConnection() *gorm.DB {
	dsnURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsnURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDatabaseConnection() *gorm.DB {

	dsnURL := os.Getenv("DATABASE_URL") + "&application_name=$ docs_simplecrud_gorm"
	db, err := gorm.Open(postgres.Open(dsnURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db

}
package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabaseConnection() *gorm.DB {
	env, err := LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dsnURL := env.DatabaseUrl + "&application_name=$ docs_simplecrud_gorm"

	db, err := gorm.Open(postgres.Open(dsnURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

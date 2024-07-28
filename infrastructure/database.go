package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDatabaseConnection() {
	dsnURL := os.Getenv("DATABASE_URL") + "&application_name=$ docs_simplecrud_gorm"

	_, err := gorm.Open(postgres.Open(dsnURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

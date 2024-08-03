package migrateService

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"stonewall-api/app/models"
)

func InitDatabaseConnection() {

	dsnURL := os.Getenv("DATABASE_URL") + "&application_name=$ docs_simplecrud_gorm"
	db, err := gorm.Open(postgres.Open(dsnURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Автомиграция
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

}

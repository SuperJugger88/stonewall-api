package migrations

import (
    "log"
	"stonewall-api/domain/models"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

// // Определение интерфейса для миграций
// type UpMigrate interface {
//     Up() error
// }

// // Структура, реализующая интерфейс UpMigrate
// type Migrator struct {
//     db *gorm.DB
// }

// // Метод Up для выполнения миграций
// func (m *Migrator) Up() error {
//     return m.db.AutoMigrate(&models.User{})
// }

func Migrate() {

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
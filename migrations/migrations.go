package migrations

import (
	"github.com/GabrielVictorP/golang-restAPI.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}

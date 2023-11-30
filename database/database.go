package database

import (
	"read-more-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Collection{}, &models.Entry{})
}

package models

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDatabase() {
	var err error
	db, err := gorm.Open(sqlite.Open(filepath.Join("_data", "test.db")), &gorm.Config{})
	if err != nil {
		log.Fatalln("[db] failed to connect database")
	}
	err = db.AutoMigrate(&Collection{}, &Entry{})
	if err != nil {
		log.Fatalln("[db] failed to migrate")
	}
	log.Println("[db] migrated")
	Database = db
}

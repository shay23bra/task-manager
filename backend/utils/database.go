package utils

import (
	"log"
	"task-manager/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.Task{})
	return db
}

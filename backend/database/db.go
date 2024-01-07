package database

import (
	"deals/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func Init() {
	db, _ = gorm.Open("sqlite3", "test.db")
	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})
}

func DeInit() {
	db.Close()
}

package database

import (
	"fmt"

	"deals/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func Init() {
	db, err := gorm.Open("sqlite3", "test2.db")
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
	}
	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})
}

func DeInit() {
	db.Close()
}

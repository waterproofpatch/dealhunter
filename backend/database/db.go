package database

import (
	"fmt"
	"log"
	"os"

	"deals/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gDb *gorm.DB

func GetDb() *gorm.DB {
	return gDb
}

func Init() {
	password := os.Getenv("DB_PASSWORD")
	host := "postgresql-waterproofpatch.alwaysdata.net"
	port := 5432
	user := "waterproofpatch"
	dbname := "waterproofpatch_deals_dev"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})
	gDb = db
}

func DeInit() {
	db, _ := gDb.DB()
	db.Close()
}

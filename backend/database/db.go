package database

import (
	"fmt"

	"deals/environment"
	"deals/logging"
	"deals/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gDb *gorm.DB

func GetDb() *gorm.DB {
	if gDb == nil {
		panic("gDb is nil!")
	}
	return gDb
}

func Init() error {
	password := environment.GetEnvironment().DB_PASSWORD
	host := environment.GetEnvironment().DB_HOST
	port := environment.GetEnvironment().DB_PORT
	user := environment.GetEnvironment().DB_USER
	dbname := environment.GetEnvironment().DB_NAME

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.GetLogger().Error().Msg(err.Error())
		return err
	}

	migrate(db)
	gDb = db

	return nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})
}

func DeInit() {
	db, _ := gDb.DB()
	db.Close()
}

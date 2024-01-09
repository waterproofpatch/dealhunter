package database

import (
	"fmt"

	"deals/environment"
	"deals/models"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gDb *gorm.DB

func Init(environment *environment.Environment, logger *zerolog.Logger) (*gorm.DB, error) {
	password := environment.DB_PASSWORD
	host := "postgresql-dealhunter.alwaysdata.net"
	port := 5432
	user := "dealhunter"
	dbname := environment.DB_NAME

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Err(err)
		return nil, err
	}
	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})
	return db, nil
}

func DeInit() {
	db, _ := gDb.DB()
	db.Close()
}

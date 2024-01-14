package database

import (
	"errors"
	"fmt"

	"deals/environment"
	"deals/logging"
	"deals/models"

	"golang.org/x/crypto/bcrypt"
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

	err = addDefaultUser(db)
	if err != nil {
		logging.GetLogger().Error().Msg("Failed adding default user")
		return errors.New("Failed adding default user")
	}
	gDb = db

	return nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})
	db.AutoMigrate(&models.User{})
}

func addDefaultUser(db *gorm.DB) error {
	var user models.User
	if err := db.Where("is_admin = ?", true).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(environment.GetEnvironment().ADMIN_PASSWORD), bcrypt.DefaultCost)
			if err != nil {
				logging.GetLogger().Error().Msgf("Failed hashing password")
				return errors.New("Failed hashing password")
			}
			defaultUser := models.User{
				Email:        environment.GetEnvironment().ADMIN_EMAIL,
				PasswordHash: string(hashedPassword),
				Reputation:   0,
				IsAdmin:      true,
			}
			if err := db.Create(&defaultUser).Error; err != nil {
				logging.GetLogger().Error().Msgf("Failed adding default user")
				return errors.New("Failed adding default user")
			}
		} else {
			logging.GetLogger().Error().Msgf("Failed retrieving user")
			return errors.New("Failed retrieving user")
		}
	}
	return nil
}

func DeInit() {
	db, _ := gDb.DB()
	db.Close()
}

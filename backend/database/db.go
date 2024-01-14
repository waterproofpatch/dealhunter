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

	setTriggers(db)
	gDb = db

	return nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Vote{})
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

func setTriggers(db *gorm.DB) {
	tables := []string{"users", "locations", "deals", "votes"}

	for _, table := range tables {
		sql := fmt.Sprintf(`
    CREATE OR REPLACE FUNCTION check_record_count() RETURNS TRIGGER AS $$
    BEGIN
       IF (SELECT count(*) FROM %s) >= 1000 THEN
          RAISE EXCEPTION 'You can store only 1000 records.';
       END IF;
       RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;

    CREATE TRIGGER your_trigger_name
    BEFORE INSERT ON %s
    FOR EACH ROW EXECUTE PROCEDURE check_record_count();
    `, table, table)

		err := db.Exec(sql).Error
		if err != nil {
			// Handle error
			logging.GetLogger().Error().Msgf("Failed executing trigger: %v", err)
		}
	}
}

func DeInit() {
	db, _ := gDb.DB()
	db.Close()
}

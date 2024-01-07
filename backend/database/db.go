package database

import "github.com/jinzhu/gorm"

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

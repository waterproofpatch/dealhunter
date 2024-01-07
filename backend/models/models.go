package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Longitude float64
	Latitude  float64
}

type Deal struct {
	gorm.Model
	Location       Location `gorm:"foreignkey:LocationID"`
	LocationID     uint
	RetailPrice    float64
	ActualPrice    float64
	StoreName      string
	ItemName       string
	Upvotes        uint
	LastUpvoteTime time.Time
}

type UserMeta struct {
	gorm.Model
	Token string
}

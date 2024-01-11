// see frontend 'models' folder
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
	Upvotes        int
	LastUpvoteTime time.Time
}

type User struct {
	gorm.Model
	Email        string `json:"-"`
	PasswordHash string `json:"-"`
	IsAdmin      bool   `json:"-"`
	Reputation   int
}

type JwtAccessToken struct {
	AccessToken string
}

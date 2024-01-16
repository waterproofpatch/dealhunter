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

type Address struct {
	gorm.Model
	Address string
}

type Deal struct {
	gorm.Model
	Location       Location `gorm:"foreignkey:LocationID"`
	LocationID     uint
	User           User `gorm:"foreignkey:UserID"` // Add this line
	UserID         uint // Add this line
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
	Deals        []Deal `gorm:"foreignkey:UserID"` // Add this line
}
type Vote struct {
	gorm.Model
	User   User `gorm:"foreignkey:UserID"`
	UserID uint
	Deal   Deal `gorm:"foreignkey:DealID"`
	DealID uint
}

type JwtAccessToken struct {
	AccessToken string
}

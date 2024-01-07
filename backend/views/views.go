package views

import (
	"encoding/json"
	"net/http"
	"time"

	"deals/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// init views handle to the db
func Init(_db *gorm.DB) {
	db = _db
}

func GetDeals(w http.ResponseWriter, r *http.Request) {
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func CreateDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	_ = json.NewDecoder(r.Body).Decode(&deal)
	db.Create(&deal)
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func UpdateDeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	vote := r.URL.Query().Get("vote")

	var deal models.Deal
	db.First(&deal, id)

	if vote == "up" {
		deal.Upvotes++
		deal.LastUpvoteTime = time.Now()
		db.Save(&deal)
	}

	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

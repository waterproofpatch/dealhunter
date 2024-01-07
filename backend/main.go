package main

import (
	"encoding/json"
	"net/http"

	"deals/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func main() {
	db, _ = gorm.Open("sqlite3", "test.db")
	defer db.Close()

	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})

	r := mux.NewRouter()
	r.HandleFunc("/deals", getDeals).Methods("GET")
	r.HandleFunc("/deals", createDeal).Methods("POST")
	http.ListenAndServe(":8000", r)
}

func getDeals(w http.ResponseWriter, r *http.Request) {
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func createDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	_ = json.NewDecoder(r.Body).Decode(&deal)
	db.Create(&deal)
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

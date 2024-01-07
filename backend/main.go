package main

import (
	"encoding/json"
	"log"
	"net/http"

	"deals/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/cors"
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
	r.HandleFunc("/deals/{id}", updateDeal).Methods("PUT")

	handler := cors.Default().Handler(r)

	serverAddress := ":8000"
	log.Printf("Server starting on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, handler))
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

func updateDeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	vote := r.URL.Query().Get("vote")

	var deal models.Deal
	db.First(&deal, id)

	if vote == "up" {
		deal.Upvotes++
		db.Save(&deal)
	}

	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

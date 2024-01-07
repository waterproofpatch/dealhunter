package main

import (
	"log"
	"net/http"

	"deals/database"
	"deals/models"
	"deals/views"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/cors"
)

func main() {
	db := database.GetDb()
	db, _ = gorm.Open("sqlite3", "test.db")
	defer db.Close()

	db.AutoMigrate(&models.Deal{})
	db.AutoMigrate(&models.Location{})

	views.Init(db)

	r := mux.NewRouter()
	r.HandleFunc("/deals", views.GetDeals).Methods("GET")
	r.HandleFunc("/deals", views.CreateDeal).Methods("POST")
	r.HandleFunc("/deals/{id}", views.UpdateDeal).Methods("PUT")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // your website
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := corsOptions.Handler(r)

	serverAddress := ":8000"
	log.Printf("Server starting on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, handler))
}

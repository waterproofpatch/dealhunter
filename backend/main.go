package main

import (
	"log"
	"net/http"

	"deals/database"
	"deals/views"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	database.Init()
	defer database.DeInit()

	db := database.GetDb()

	handler := views.Init(db)

	serverAddress := ":8000"
	log.Printf("Server starting on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, *handler))
}

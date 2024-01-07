package main

import (
	"log"
	"net/http"
	"os"

	"deals/database"
	"deals/views"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	database.Init()
	defer database.DeInit()

	db := database.GetDb()

	handler, _ := views.Init(db)
	// authentication.Init("someSecret",
	// 	"someRefreshSecret",
	// 	"defaultAdminEmail",
	// 	"defaultAdminUsername",
	// 	"defaultAdminPassword",
	// 	router,
	// 	"dbUrl",
	// 	false,
	// 	false,
	// 	nil,
	// 	"")

	// serverAddress := ":8000"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Provide a default value if no environment variable is set
	}
	serverAddress := ":" + port
	log.Printf("Server starting on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, *handler))
}

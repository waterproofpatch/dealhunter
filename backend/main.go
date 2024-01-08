package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"deals/database"
	"deals/views"
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
		fmt.Println("No port environment variable set, defaulting to 8000")
		port = "8000" // Provide a default value if no environment variable is set
	}
	serverAddress := ":" + port
	log.Printf("Server starting on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, *handler))
}

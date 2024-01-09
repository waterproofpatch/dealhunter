package main

import (
	"fmt"
	"log"
	"net/http"

	"deals/database"
	"deals/environment"
	"deals/logging"
	"deals/views"
)

func main() {
	// init logging library first
	logger := logging.Init()

	// environment init
	env, err := environment.Init()
	if err != nil {
		logger.Err(err)
		return
	}

	// init database library next
	db, err := database.Init(env, logger)
	if err != nil {
		logger.Err(err)
		return
	}
	defer database.DeInit()

	// init views with the db
	handler, _ := views.Init(env, db, logger)

	port := env.PORT
	if port == "" {
		fmt.Println("No port environment variable set, defaulting to 8000")
		port = "8000" // Provide a default value if no environment variable is set
	}

	serverAddress := ":" + port
	log.Printf("Server starting on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, *handler))
}

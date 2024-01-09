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
	logging.Init()

	// environment init
	err := environment.Init()
	if err != nil {
		logging.GetLogger().Error().Msg(err.Error())
		return
	}

	// init database library next
	err = database.Init()
	if err != nil {
		logging.GetLogger().Error().Msg(err.Error())
		return
	}
	defer database.DeInit()

	// init views with the db
	handler, _ := views.Init()

	port := environment.GetEnvironment().PORT
	if port == "" {
		fmt.Println("No port environment variable set, defaulting to 8000")
		port = "8000" // Provide a default value if no environment variable is set
	}

	serverAddress := ":" + port
	log.Printf("Server starting on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, *handler))
}

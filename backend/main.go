package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"deals/database"
	"deals/environment"
	"deals/logging"
	"deals/views"
)

// curl me like:
// curl -X POST http://localhost:8000/shutdown
func shutdown(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Shutdown request received. Shutting down...")
	os.Exit(0)
}

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
	handler, r := views.Init()
	// bind shutdown handler
	r.HandleFunc("/shutdown", shutdown)

	port := environment.GetEnvironment().PORT
	if port == "" {
		fmt.Println("No port environment variable set, defaulting to 8000")
		port = "8000" // Provide a default value if no environment variable is set
	}

	serverAddress := ":" + port
	log.Printf("Server starting on %s\n", serverAddress)
	// log.Fatal(http.ListenAndServe(serverAddress, *handler))
	// log.Fatal(http.ListenAndServeTLS(serverAddress, "server.crt", "server.key", *handler))
	if os.Getenv("LOCAL_DEV") != "" {
		log.Println("Starting server in production mode (expect load balanced TLS)...")
		log.Fatal(http.ListenAndServeTLS(serverAddress, "server.crt", "server.key", *handler))
	} else {
		log.Println("Starting server in local development mode (local TLS)...")
		log.Fatal(http.ListenAndServe(serverAddress, *handler))
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"deals/database"
	"deals/environment"
	"deals/logging"
	"deals/views"
)

func getLocation() {
	lat := "37.4224764"
	lng := "-122.0842499"
	apiKey := environment.GetEnvironment().GOOGLE_GEOCODING_API_KEY

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", lat, lng, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
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
	handler, _ := views.Init()

	port := environment.GetEnvironment().PORT
	if port == "" {
		fmt.Println("No port environment variable set, defaulting to 8000")
		port = "8000" // Provide a default value if no environment variable is set
	}
	getLocation()

	serverAddress := ":" + port
	log.Printf("Server starting on %s\n", serverAddress)
	// log.Fatal(http.ListenAndServe(serverAddress, *handler))
	log.Fatal(http.ListenAndServeTLS(serverAddress, "server.crt", "server.key", *handler))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"deals/database"
	"deals/environment"
	"deals/logging"
	"deals/views"
)

type Location struct {
	Address string `json:"formatted_address"`
}

func getLocationFor(lat string, lon string) string {
	apiKey := environment.GetEnvironment().GOOGLE_GEOCODING_API_KEY

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", lat, lon, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Write the response body to a text file
	err = ioutil.WriteFile("response.txt", body, 0o644)
	if err != nil {
		panic(err)
	}

	var location Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		panic(err)
	}

	return location.Address
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
	address := getLocationFor("39.148", "-76.73")
	fmt.Printf("address is %v", address)

	serverAddress := ":" + port
	log.Printf("Server starting on %s\n", serverAddress)
	// log.Fatal(http.ListenAndServe(serverAddress, *handler))
	log.Fatal(http.ListenAndServeTLS(serverAddress, "server.crt", "server.key", *handler))
}

package location

import (
	"context"
	"fmt"
	"log"

	"deals/environment"

	"googlemaps.github.io/maps"
)

func GetLatLonFor(address string) (float64, float64) {
	// Create a new client with your API key
	ctx := context.Background()
	// Create a context with a cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // cancel the context when the main function returns

	c, err := maps.NewClient(maps.WithAPIKey(environment.GetEnvironment().GOOGLE_GEOCODING_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// Reverse geocode a latitude and longitude
	r := &maps.GeocodingRequest{
		Address: address,
	}
	resp, err := c.Geocode(ctx, r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// Check if the response is not empty
	if len(resp) > 0 {
		lat := resp[0].Geometry.Location.Lat
		lng := resp[0].Geometry.Location.Lng
		return lat, lng
	}

	// Return zero values if no result was found
	return 0, 0
}

func GetLocationFor(lat float64, lon float64) string {
	// Create a new client with your API key
	ctx := context.Background()
	// Create a context with a cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // cancel the context when the main function returns

	c, err := maps.NewClient(maps.WithAPIKey(environment.GetEnvironment().GOOGLE_GEOCODING_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// Reverse geocode a latitude and longitude
	r := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lon,
		},
	}
	resp, err := c.ReverseGeocode(ctx, r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// Print the results
	firstAddress := ""
	for _, result := range resp {
		fmt.Println(result.FormattedAddress)
		if firstAddress == "" {
			firstAddress = result.FormattedAddress
		}
	}
	return firstAddress
}

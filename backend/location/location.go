package location

import (
	"context"
	"fmt"
	"log"

	"deals/environment"

	"googlemaps.github.io/maps"
)

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

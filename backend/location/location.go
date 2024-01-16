package location

import (
	"context"
	"fmt"
	"log"

	"deals/database"
	"deals/environment"
	"deals/logging"
	"deals/models"

	"github.com/jinzhu/gorm"
	"googlemaps.github.io/maps"
)

// try and get a cached location so we don't query google
// return "" on cache miss
func getAddressForFromCache(lat float64, lon float64) string {
	var addressCache models.AddressCache
	if err := database.GetDb().First(&addressCache, "latitude = ? AND longitude = ?", lat, lon).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ""
		}
		logging.GetLogger().Debug().Msgf("Error querying the database: %v", err)
		return ""
	}
	return addressCache.Address
}

// try and get a cached lat, lon so we don't query google.
// return 0, 0 on cache miss.
func getLatLonForFromCache(address string) (float64, float64) {
	var addressCache models.AddressCache
	if err := database.GetDb().First(&addressCache, "address = ?", address).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, 0
		}
		logging.GetLogger().Debug().Msgf("Error querying the database: %v", err)
		return 0, 0
	}
	return addressCache.Latitude, addressCache.Longitude
}

func GetLatLonFor(address string) (float64, float64) {
	// temp quota
	lat, lon := getLatLonForFromCache(address)
	if lat != -1 && lon != -1 {
		logging.GetLogger().Debug().Msgf("returning lat, lon %v, %v from cache.", lat, lon)
		return lat, lon
	}

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
		// possible quota issue
		logging.GetLogger().Debug().Msgf("fatal error: %s", err)
		return 0, 0
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

func GetAddressFor(lat float64, lon float64) string {
	cachedAddress := getAddressForFromCache(lat, lon)
	if cachedAddress != "" {
		logging.GetLogger().Debug().Msgf("returning location %v from cache.", cachedAddress)
		return cachedAddress
	}
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
		logging.GetLogger().Debug().Msgf("fatal error: %s", err)
		return "over_quota"
	}

	// TODO someday return all candidates and let the user choose
	firstAddress := ""
	for _, result := range resp {
		fmt.Println(result.FormattedAddress)
		if firstAddress == "" {
			firstAddress = result.FormattedAddress
		}
	}
	return firstAddress
}

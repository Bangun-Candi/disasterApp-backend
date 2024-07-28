package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"math"
	"net/http"
	"strconv"
	"users/models"
	"users/utils"

	"github.com/go-resty/resty/v2"
)

type DisasterStatus struct {
	LocationName       string `json:"locationName"`
	StatusLocation     string `json:"statusLocation"`
	StatusDisasterCode string `json:"statusDisasterCode"`
	SendRescueStatus   string `json:"sendRescueStatus"`
	DisasterName       string `json:"disasterName"`
	DisasterNameCode   string `json:"disasterNameCode"`
}

type NominatimResponse struct {
	DisplayName string `json:"display_name"`
}

func GetCurrentStatus(userEmail, longitudeLocation, latitudeLocation string) (DisasterStatus, error) {
	db := utils.GetDB()

	// Validate if the user exists
	var user models.User
	err := db.QueryRow("SELECT id FROM users WHERE email = ?", userEmail).Scan(&user.ID)
	if err == sql.ErrNoRows {
		return DisasterStatus{}, errors.New("user not found")
	} else if err != nil {
		return DisasterStatus{}, err
	}

	// Convert latitude and longitude to float64
	lat, err := strconv.ParseFloat(latitudeLocation, 64)
	if err != nil {
		return DisasterStatus{}, errors.New("invalid latitude")
	}
	long, err := strconv.ParseFloat(longitudeLocation, 64)
	if err != nil {
		return DisasterStatus{}, errors.New("invalid longitude")
	}

	// Get location name from Nominatim API
	locationName, err := getLocationName(lat, long)
	if err != nil {
		return DisasterStatus{}, err
	}

	// For the purpose of this example, we'll use hardcoded values for disaster status
	// In a real application, you would fetch this data from a database or external service
	distanceToDisaster := calculateDistance(lat, long, -6.200000, 106.816666) // Example: distance to a disaster location
	statusDisasterCode := "03"                                                // Default to safe
	if distanceToDisaster < 50 {
		statusDisasterCode = "01" // Danger if within 50 km
	} else if distanceToDisaster < 100 {
		statusDisasterCode = "02" // Warning if within 100 km
	}

	sendRescueStatus := "FALSE"
	if statusDisasterCode == "01" {
		sendRescueStatus = "TRUE"
	}

	disasterStatus := DisasterStatus{
		LocationName:       locationName,
		StatusLocation:     "Active", // Replace with actual status
		StatusDisasterCode: statusDisasterCode,
		SendRescueStatus:   sendRescueStatus,
		DisasterName:       "Gempa Bumi",
		DisasterNameCode:   "001",
	}

	return disasterStatus, nil
}

func calculateDistance(lat1, long1, lat2, long2 float64) float64 {
	// Haversine formula to calculate the distance between two points
	const R = 6371 // Radius of the Earth in km
	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLong := (long2 - long1) * (math.Pi / 180.0)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180.0))*math.Cos(lat2*(math.Pi/180.0))*
			math.Sin(dLong/2)*math.Sin(dLong/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c
	return distance
}

func SendRealtimeLocation(request models.RealTimeLocation) (models.RealTimeLocation, error) {
	db := utils.GetDB()

	// Get location name from Nominatim API
	locationName, err := getLocationName(request.Latitude, request.Longitude)
	if err != nil {
		return models.RealTimeLocation{}, err
	}
	request.LocationName = locationName

	// Insert or update real-time location into the database
	query := `
		INSERT INTO real_time_locations (user_id, latitude, longitude, location_name, disaster_code, disaster_name)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		latitude = VALUES(latitude),
		longitude = VALUES(longitude),
		location_name = VALUES(location_name),
		disaster_code = VALUES(disaster_code),
		disaster_name = VALUES(disaster_name)
	`
	_, err = db.Exec(query, request.UserID, request.Latitude, request.Longitude, request.LocationName, request.DisasterCode, request.DisasterName)
	if err != nil {
		return models.RealTimeLocation{}, err
	}

	return request, nil
}

func getLocationName(lat, long float64) (string, error) {
	client := resty.New()
	nominatimURL := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", lat, long)

	resp, err := client.R().Get(nominatimURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New("failed to fetch location name from Nominatim")
	}

	var nominatimResponse struct {
		DisplayName string `json:"display_name"`
	}
	err = json.Unmarshal(resp.Body(), &nominatimResponse)
	if err != nil {
		return "", err
	}

	return nominatimResponse.DisplayName, nil
}

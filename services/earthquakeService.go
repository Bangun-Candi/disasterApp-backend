package services

import (
	"encoding/xml"
	"errors"
	"net/http"
	"strconv"
	"time"
	"users/models"

	"github.com/go-resty/resty/v2"
)

const bmkgURL = "https://data.bmkg.go.id/DataMKG/TEWS/autogempa.xml"

// EarthquakeResponse represents the BMKG earthquake response
type EarthquakeResponse struct {
	XMLName    xml.Name `xml:"Infogempa"`
	Earthquake struct {
		XMLName     xml.Name `xml:"gempa"`
		DateTime    string   `xml:"DateTime"`
		Coordinates string   `xml:"Coordinates"`
		Latitude    string   `xml:"Lintang"`
		Longitude   string   `xml:"Bujur"`
		Magnitude   string   `xml:"Magnitude"`
		Depth       string   `xml:"Kedalaman"`
		Region      string   `xml:"Wilayah"`
		Potential   string   `xml:"Potensi"`
		Shakemap    string   `xml:"Shakemap"`
	} `xml:"gempa"`
}

func FetchBMKGEarthquakeData() (models.Earthquake, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/xml").
		Get(bmkgURL)

	if err != nil {
		return models.Earthquake{}, err
	}

	if resp.StatusCode() != http.StatusOK {
		return models.Earthquake{}, errors.New("failed to fetch data from BMKG")
	}

	var bmkgResponse EarthquakeResponse
	err = xml.Unmarshal(resp.Body(), &bmkgResponse)
	if err != nil {
		return models.Earthquake{}, err
	}

	earthquake, err := mapBMKGEarthquakeData(bmkgResponse)
	if err != nil {
		return models.Earthquake{}, err
	}

	return earthquake, nil
}

func mapBMKGEarthquakeData(bmkgResponse EarthquakeResponse) (models.Earthquake, error) {
	layout := "2006-01-02T15:04:05-07:00"
	occurrenceTime, err := time.Parse(layout, bmkgResponse.Earthquake.DateTime)
	if err != nil {
		return models.Earthquake{}, err
	}

	magnitude := parseStringToFloat(bmkgResponse.Earthquake.Magnitude)

	return models.Earthquake{
		Latitude:       parseStringToFloat(bmkgResponse.Earthquake.Latitude),
		Longitude:      parseStringToFloat(bmkgResponse.Earthquake.Longitude),
		Magnitude:      magnitude,
		Depth:          parseStringToFloat(bmkgResponse.Earthquake.Depth),
		Location:       bmkgResponse.Earthquake.Region,
		OccurrenceTime: occurrenceTime,
		Status:         determineStatus(magnitude),
	}, nil
}

func parseStringToFloat(input string) float64 {
	value, _ := strconv.ParseFloat(input, 64)
	return value
}

func determineStatus(magnitude float64) string {
	if magnitude == 0.0 {
		return "safe"
	} else if magnitude < 5.0 {
		return "normal"
	}
	return "danger"
}

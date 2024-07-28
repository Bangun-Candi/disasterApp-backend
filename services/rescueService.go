package services

import (
	"database/sql"
	"errors"
	"users/models"
	"users/utils"
)

type RescuersCategory struct {
	CategoryName string `json:"categoryName"`
	CategoryCode string `json:"categoryCode"`
}

type RescueRequest struct {
	UserEmail          string        `json:"userEmail"`
	UserName           string        `json:"userName"`
	UserPhoneNumber    string        `json:"userPhoneNumber"`
	LongitudeLocation  string        `json:"longitudeLocation"`
	LatitudeLocation   string        `json:"latitudeLocation"`
	RescuersCode       []RescuerCode `json:"rescuersCode"`
	DisasterTypeCode   string        `json:"disasterTypeCode"`
	StatusDisaster     string        `json:"statusDisaster"`
	StatusDisasterCode string        `json:"statusDisasterCode"`
	Notes              string        `json:"notes"`
}

type RescuerCode struct {
	Code string `json:"code"`
}

type RescuersResponse struct {
	CategoryName string `json:"categoryName"`
	CategoryCode string `json:"categoryCode"`
}

func SendRescueDisaster(request models.RescueRequest) ([]models.Rescuer, error) {
	db := utils.GetDB()

	// Insert disaster details into the database
	query := "INSERT INTO disasters (disaster_name, disaster_code, status_code, latitude, longitude, notes) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, request.DisasterTypeCode, request.StatusDisasterCode, request.StatusDisaster, request.LatitudeLocation, request.LongitudeLocation, request.Notes)
	if err != nil {
		return nil, err
	}

	// Retrieve rescuer details based on provided codes
	var rescuers []models.Rescuer
	for _, code := range request.RescuersCode {
		var rescuer models.Rescuer
		query := "SELECT category_name, category_code FROM rescuers WHERE category_code = ?"
		err := db.QueryRow(query, code.Code).Scan(&rescuer.CategoryName, &rescuer.CategoryCode)
		if err == sql.ErrNoRows {
			return nil, errors.New("no rescuer found for code " + code.Code)
		} else if err != nil {
			return nil, err
		}
		rescuers = append(rescuers, rescuer)
	}
	return rescuers, nil
}

package services

import "users/utils"
import "users/models"

type RescuerCategory struct {
	CategoryName string `json:"categoryName"`
	CategoryCode string `json:"categoryCode"`
}

func GetRescuersCategory() ([]models.Rescuer, error) {
	db := utils.GetDB()
	var rescuers []models.Rescuer
	query := "SELECT category_name, category_code FROM rescuers"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rescuer models.Rescuer
		err := rows.Scan(&rescuer.CategoryName, &rescuer.CategoryCode)
		if err != nil {
			return nil, err
		}
		rescuers = append(rescuers, rescuer)
	}
	return rescuers, nil
}

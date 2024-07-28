package services

import (
	"users/models"
	"users/utils"
)

func GetInvestmentReferences() ([]models.InvestmentReference, error) {
	db := utils.GetDB()
	var investments []models.InvestmentReference
	query := "SELECT id, type, name, description FROM investment_references"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var investment models.InvestmentReference
		err := rows.Scan(&investment.ID, &investment.Type, &investment.Name, &investment.Description)
		if err != nil {
			return nil, err
		}
		investments = append(investments, investment)
	}
	return investments, nil
}

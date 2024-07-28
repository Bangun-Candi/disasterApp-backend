package services

import (
	"users/models"
	"users/utils"
)

func GetVentureCapital() ([]models.VentureCapital, error) {
	db := utils.GetDB()
	var vcs []models.VentureCapital
	query := "SELECT id, name, contact_info, address, investment_size, created_at FROM venture_capital"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var vc models.VentureCapital
		err := rows.Scan(&vc.ID, &vc.Name, &vc.ContactInfo, &vc.Address, &vc.InvestmentSize, &vc.CreatedAt)
		if err != nil {
			return nil, err
		}
		vcs = append(vcs, vc)
	}
	return vcs, nil
}

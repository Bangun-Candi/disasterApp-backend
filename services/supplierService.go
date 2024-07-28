package services

import (
	"users/models"
	"users/utils"
)

func GetSuppliers() ([]models.Supplier, error) {
	db := utils.GetDB()
	var suppliers []models.Supplier
	query := "SELECT id, name, contact_info, category, created_at FROM suppliers"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier models.Supplier
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.ContactInfo, &supplier.Category, &supplier.CreatedAt)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers, nil
}

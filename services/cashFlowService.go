package services

import (
	"time"
	"users/models"
	"users/utils"
)

func GetCashFlowReport(userID int, startDate, endDate time.Time) ([]models.CashFlow, error) {
	db := utils.GetDB()
	var cashFlows []models.CashFlow
	query := `SELECT id, user_id, amount, type, description, created_at 
	          FROM cash_flow 
	          WHERE user_id = ? AND created_at BETWEEN ? AND ?`
	rows, err := db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cashFlow models.CashFlow
		err := rows.Scan(&cashFlow.ID, &cashFlow.UserID, &cashFlow.Amount, &cashFlow.Type, &cashFlow.Description, &cashFlow.CreatedAt)
		if err != nil {
			return nil, err
		}
		cashFlows = append(cashFlows, cashFlow)
	}
	return cashFlows, nil
}

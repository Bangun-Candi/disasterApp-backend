package services

import (
	"database/sql"
	"errors"
	"time"
	"users/models"
	"users/utils"
)

func GetBalance(userID, userEmail string) (models.Wallet, error) {
	db := utils.GetDB()
	var wallet models.Wallet
	query := "SELECT id, user_id, balance FROM wallets WHERE user_id = ?"
	err := db.QueryRow(query, userID).Scan(&wallet.ID, &wallet.UserID, &wallet.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return wallet, errors.New("no balance data found")
		}
		return wallet, err
	}
	return wallet, nil
}

func GetBalanceHistory(userID, userEmail string, startDate, endDate time.Time) ([]models.BalanceHistory, error) {
	db := utils.GetDB()
	var history []models.BalanceHistory
	query := `SELECT balance, status, from_account_name, from_account_number, to_account_name, to_account_number, date 
	          FROM balance_history 
	          WHERE user_id = ? AND email = ? AND date BETWEEN ? AND ?`
	rows, err := db.Query(query, userID, userEmail, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bh models.BalanceHistory
		err := rows.Scan(&bh.Balance, &bh.Status, &bh.FromAccountName, &bh.FromAccountNumber, &bh.ToAccountName, &bh.ToAccountNumber, &bh.Date)
		if err != nil {
			return nil, err
		}
		history = append(history, bh)
	}
	return history, nil
}

func GetCompanyGrowth(userID, userEmail string, startDate, endDate time.Time) (models.CompanyGrowth, error) {
	db := utils.GetDB()
	var growth models.CompanyGrowth
	query := `SELECT date, remaining_balance 
	          FROM company_growth 
	          WHERE user_id = ? AND email = ? AND date BETWEEN ? AND ?`
	err := db.QueryRow(query, userID, userEmail, startDate, endDate).Scan(&growth.Date, &growth.RemainingBalance)
	return growth, err
}

func GetSalesGrowth(userID, userEmail string, startDate, endDate time.Time) ([]models.SalesGrowth, error) {
	db := utils.GetDB()
	var sales []models.SalesGrowth
	query := `SELECT date, total_transaction 
	          FROM sales_growth 
	          WHERE user_id = ? AND email = ? AND date BETWEEN ? AND ?`
	rows, err := db.Query(query, userID, userEmail, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sg models.SalesGrowth
		err := rows.Scan(&sg.Date, &sg.TotalTransaction)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sg)
	}
	return sales, nil
}

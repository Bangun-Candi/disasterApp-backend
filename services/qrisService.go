package services

import (
	"database/sql"
	"errors"
	"users/models"
	"users/utils"
)

func GenerateQRCode(userID int, amount int64) (models.QRISTransaction, error) {
	db := utils.GetDB()
	var qrisTransaction models.QRISTransaction

	// Generate QR code (For example purpose, using a simple string. Replace with actual QR code generation)
	qrCode := "QR_CODE_" + string(userID) + "_" + string(amount)

	query := "INSERT INTO qris_transactions (user_id, amount, qr_code, status) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, userID, amount, qrCode, "PENDING")
	if err != nil {
		return qrisTransaction, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return qrisTransaction, err
	}

	qrisTransaction = models.QRISTransaction{
		ID:        int(id),
		UserID:    userID,
		Amount:    amount,
		QRCode:    qrCode,
		Status:    "PENDING",
		// CreatedAt: "null",
		// UpdatedAt: "null",
	}
	return qrisTransaction, nil
}

func ConfirmPayment(qrCode string) (models.QRISTransaction, error) {
	db := utils.GetDB()
	var qrisTransaction models.QRISTransaction

	query := "SELECT id, user_id, amount, qr_code, status, created_at, updated_at FROM qris_transactions WHERE qr_code = ?"
	err := db.QueryRow(query, qrCode).Scan(&qrisTransaction.ID, &qrisTransaction.UserID, &qrisTransaction.Amount, &qrisTransaction.QRCode, &qrisTransaction.Status, &qrisTransaction.CreatedAt, &qrisTransaction.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return qrisTransaction, errors.New("QR code not found")
		}
		return qrisTransaction, err
	}

	if qrisTransaction.Status != "PENDING" {
		return qrisTransaction, errors.New("QR code is not pending")
	}

	updateQuery := "UPDATE qris_transactions SET status = ? WHERE qr_code = ?"
	_, err = db.Exec(updateQuery, "COMPLETED", qrCode)
	if err != nil {
		return qrisTransaction, err
	}

	qrisTransaction.Status = "COMPLETED"
	return qrisTransaction, nil
}

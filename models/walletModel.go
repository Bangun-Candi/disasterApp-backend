package models

import "time"

type Wallet struct {
	ID      int    `json:"id"`
	UserID  string `json:"user_id"`
	Balance int64  `json:"balance"`
}

type BalanceHistory struct {
	Balance           int64  `json:"balance"`
	Status            string `json:"status"`
	FromAccountName   string `json:"from_account_name"`
	FromAccountNumber string `json:"from_account_number"`
	ToAccountName     string `json:"to_account_name"`
	ToAccountNumber   string `json:"to_account_number"`
	Date              string `json:"date"`
}

type CompanyGrowth struct {
	Date             string `json:"date"`
	RemainingBalance int64  `json:"remaining_balance"`
}

type SalesGrowth struct {
	Date             string `json:"date"`
	TotalTransaction int    `json:"total_transaction"`
}

type QRISTransaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Amount    int64     `json:"amount"`
	QRCode    string    `json:"qr_code"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CashFlow struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Amount      int64     `json:"amount"`
	Type        string    `json:"type"` // IN or OUT
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type InvestmentReference struct {
	ID          int       `json:"id"`
	Type        string    `json:"type"` // high_risk or low_risk
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Supplier struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	ContactInfo string    `json:"contact_info"`
	Category   string    `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
}

type VentureCapital struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	ContactInfo   string    `json:"contact_info"`
	Address       string    `json:"address"`
	InvestmentSize float64  `json:"investment_size"`
	CreatedAt     time.Time `json:"created_at"`
}

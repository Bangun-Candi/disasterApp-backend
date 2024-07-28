package services

import (
	"fmt"
	"math/rand"
	"time"
	"users/models"
	"users/utils"
)

// func CreateUser(user *models.User) error {
// 	query := "INSERT INTO users (email, username, phone_number, password_hash, pin_hash, kyc_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())"
// 	stmt, err := utils.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(user.Email, user.Username, user.PhoneNumber, user.PasswordHash, user.PinHash, user.KycStatus)
// 	return err
// }

func SaveOTP(otp *models.OTPCode) error {
	query := "INSERT INTO otp_codes (email, phone_number, otp_code, type, expires_at, created_at) VALUES (?, ?, ?, ?, ?, NOW())"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(otp.Email, otp.PhoneNumber, otp.OTPCode, otp.Type, otp.ExpiresAt)
	return err
}

func VerifyOTP(email, phoneNumber, otpCode, otpType string) (bool, error) {
	var otp models.OTPCode
	var expiresAtStr, createdAtStr []uint8
	query := "SELECT id, email, phone_number, otp_code, type, expires_at, created_at FROM otp_codes WHERE email = ? AND phone_number = ? AND otp_code = ? AND type = ?"
	row := utils.DB.QueryRow(query, email, phoneNumber, otpCode, otpType)
	err := row.Scan(&otp.ID, &otp.Email, &otp.PhoneNumber, &otp.OTPCode, &otp.Type, &expiresAtStr, &createdAtStr)
	if err != nil {
		fmt.Println("Error querying OTP:", err)
		return false, err
	}

	// Convert expiresAtStr and createdAtStr to time.Time
	expiresAt, err := time.Parse("2006-01-02 15:04:05", string(expiresAtStr))
	if err != nil {
		fmt.Println("Error parsing expires_at:", err)
		return false, err
	}
	otp.ExpiresAt = expiresAt

	createdAt, err := time.Parse("2006-01-02 15:04:05", string(createdAtStr))
	if err != nil {
		fmt.Println("Error parsing created_at:", err)
		return false, err
	}
	otp.CreatedAt = createdAt

	fmt.Println("OTP from database:", otp.OTPCode)
	fmt.Println("OTP from request:", otpCode)
	if time.Now().After(otp.ExpiresAt) {
		fmt.Println("OTP has expired")
		return false, nil
	}

	return otp.OTPCode == otpCode, nil
}

// func SaveKYCData(kycData *models.KYCData) error {
// 	query := "INSERT INTO kyc_data (user_id, nik, name, birthday_date, birthday_place, gender, address, religion, marriage_status, occupation, citizenship, validity_period, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())"
// 	stmt, err := utils.DB.Prepare(query)
// 	if err != nil {
// 		log.Println("Error preparing query:", err)
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(
// 		kycData.UserID,
// 		kycData.NIK,
// 		kycData.Name,
// 		kycData.BirthdayDate,
// 		kycData.BirthdayPlace,
// 		kycData.Gender,
// 		kycData.Address,
// 		kycData.Religion,
// 		kycData.MarriageStatus,
// 		kycData.Occupation,
// 		kycData.Citizenship,
// 		kycData.ValidityPeriod,
// 	)
// 	if err != nil {
// 		log.Println("Error executing query:", err)
// 		return err
// 	}

// 	log.Println("KYC data saved successfully")
// 	return nil
// }
// func GetUserByEmail(email string) (*models.User, error) {
// 	var user models.User
// 	query := "SELECT id, email, username, phone_number, password_hash, pin_hash, kyc_status, created_at, updated_at FROM users WHERE email = ?"
// 	err := utils.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Username, &user.PhoneNumber, &user.PasswordHash, &user.PinHash, &user.KycStatus, &user.CreatedAt, &user.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func GetUserByID(id int) (*models.User, error) {
// 	var user models.User
// 	var createdAtStr, updatedAtStr []uint8
// 	query := "SELECT id, email, username, phone_number, password_hash, pin_hash, kyc_status, created_at, updated_at FROM users WHERE id = ?"
// 	err := utils.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Username, &user.PhoneNumber, &user.PasswordHash, &user.PinHash, &user.KycStatus, &createdAtStr, &updatedAtStr)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("user with ID %d not found", id)
// 		}
// 		return nil, err
// 	}

// 	// Convert createdAtStr and updatedAtStr to time.Time
// 	createdAt, err := time.Parse("2006-01-02 15:04:05", string(createdAtStr))
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing created_at: %v", err)
// 	}
// 	user.CreatedAt = createdAt

// 	updatedAt, err := time.Parse("2006-01-02 15:04:05", string(updatedAtStr))
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing updated_at: %v", err)
// 	}
// 	user.UpdatedAt = updatedAt

// 	return &user, nil
// }

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func CheckPinHash(pin, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pin))
// 	return err == nil
// }

func GenerateOTPCode() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// func ValidateKYCData(kycData *models.KYCData) error {
// 	var (
// 		existingKYC    models.KYCData
// 		birthdayDate   []uint8
// 		validityPeriod []uint8
// 	)
// 	query := "SELECT nik, name, birthday_date, birthday_place, gender, address, religion, marriage_status, occupation, citizenship, validity_period FROM kyc_data WHERE nik = ?"
// 	err := utils.DB.QueryRow(query, kycData.NIK).Scan(
// 		&existingKYC.NIK,
// 		&existingKYC.Name,
// 		&birthdayDate,
// 		&existingKYC.BirthdayPlace,
// 		&existingKYC.Gender,
// 		&existingKYC.Address,
// 		&existingKYC.Religion,
// 		&existingKYC.MarriageStatus,
// 		&existingKYC.Occupation,
// 		&existingKYC.Citizenship,
// 		&validityPeriod,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			log.Println("KYC data not found for NIK:", kycData.NIK)
// 			return errors.New("KYC data not found")
// 		}
// 		log.Println("Error querying KYC data:", err)
// 		return errors.New("failed to validate KYC data")
// 	}

// 	// Convert the birthdayDate and validityPeriod to time.Time
// 	existingKYC.BirthdayDate, err = time.Parse("2006-01-02", string(birthdayDate))
// 	if err != nil {
// 		log.Println("Error parsing birthday_date:", err)
// 		return errors.New("failed to parse birthday_date")
// 	}

// 	existingKYC.ValidityPeriod, err = time.Parse("2006-01-02", string(validityPeriod))
// 	if err != nil {
// 		log.Println("Error parsing validity_period:", err)
// 		return errors.New("failed to parse validity_period")
// 	}

// 	// Log the retrieved KYC data
// 	log.Println("Retrieved KYC data:", existingKYC)

// 	// Additional validation logic can be added here
// 	if existingKYC.NIK != kycData.NIK {
// 		log.Println("Mismatch in NIK")
// 		return errors.New("KYC data mismatch")
// 	}
// 	if existingKYC.Name != kycData.Name {
// 		log.Println("Mismatch in Name")
// 		return errors.New("KYC data mismatch")
// 	}
// 	// Add more validation as needed

// 	return nil
// }

func CheckEmail(email string) (bool, error) {
	db := utils.GetDB()
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)"
	err := db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

// func RegisterUser(user models.User) error {
// 	db := utils.GetDB()
// 	hashedPassword, _ := utils.HashPassword(user.Password)
// 	query := `
// 		INSERT INTO users (email, phone_number, username, nik, name, birthday_date, birthday_place, gender, address, religion, marriage_status, occupation, citizenship, validity_period, password, pin)
// 		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
// 	_, err := db.Exec(query, user.Email, user.PhoneNumber, user.Username, user.NIK, user.Name, user.BirthdayDate, user.BirthdayPlace, user.Gender, user.Address, user.Religion, user.MarriageStatus, user.Occupation, user.Citizenship, user.ValidityPeriod, hashedPassword, user.PIN)
// 	return err
// }

// func AuthenticateUser(email, password string) (*models.User, error) {
// 	db := utils.GetDB()
// 	var user models.User
// 	query := "SELECT id, email, phone_number, username, password FROM users WHERE email = ?"
// 	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PhoneNumber, &user.Username, &user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if utils.CheckPasswordHash(password, user.Password) {
// 		return &user, nil
// 	}
// 	return nil, nil
// }

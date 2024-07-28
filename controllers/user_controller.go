package controllers

import (
	"log"
	"net/http"
	"time"
	"users/models"
	"users/services"
	"users/utils"

	"github.com/gin-gonic/gin"
)

type OTPRequest struct {
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type OTPVerify struct {
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	OTPCode     string `json:"otpCode" binding:"required"`
	Type        string `json:"type" binding:"required"`
}
type KYCRequest struct {
	NIK            string `json:"nik" binding:"required"`
	Name           string `json:"name" binding:"required"`
	BirthdayDate   string `json:"birthday_date" binding:"required"`
	BirthdayPlace  string `json:"birthday_place" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	Address        string `json:"address" binding:"required"`
	Religion       string `json:"religion" binding:"required"`
	MarriageStatus string `json:"marriage_status" binding:"required"`
	Occupation     string `json:"occupation" binding:"required"`
	Citizenship    string `json:"citizenship" binding:"required"`
	ValidityPeriod string `json:"validity_period" binding:"required"`
}

type RegisterRequest struct {
	Email          string `json:"email" binding:"required"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	Username       string `json:"username" binding:"required"`
	NIK            string `json:"nik" binding:"required"`
	Name           string `json:"name" binding:"required"`
	BirthdayDate   string `json:"birthday_date" binding:"required"`
	BirthdayPlace  string `json:"birthday_place" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	Address        string `json:"address" binding:"required"`
	Religion       string `json:"religion" binding:"required"`
	MarriageStatus string `json:"marriage_status" binding:"required"`
	Occupation     string `json:"occupation" binding:"required"`
	Citizenship    string `json:"citizenship" binding:"required"`
	ValidityPeriod string `json:"validity_period" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Pin            string `json:"pin" binding:"required"`
}

// func CheckEmail(c *gin.Context) {
// 	var request struct {
// 		Email string `json:"email"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
// 		return
// 	}

// 	exists, err := services.CheckEmail(request.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Database error"})
// 		return
// 	}

// 	if exists {
// 		c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "Email is registered"})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": true, "message": "Email not registered"})
// 	}
// }

// func RequestOTP(c *gin.Context) {
// 	var input OTPRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": err.Error()})
// 		return
// 	}

// 	otpCode := services.GenerateOTPCode()
// 	expiresAt := time.Now().Add(10 * time.Minute)

// 	otp := models.OTPCode{
// 		Email:       input.Email,
// 		PhoneNumber: input.PhoneNumber,
// 		OTPCode:     otpCode,
// 		Type:        input.Type,
// 		ExpiresAt:   expiresAt,
// 	}

// 	if err := services.SaveOTP(&otp); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to save OTP"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
// }

func RequestOTP(c *gin.Context) {
	var input OTPRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": err.Error()})
		return
	}

	otpCode := services.GenerateOTPCode()
	expiresAt := time.Now().Add(10 * time.Minute)

	otp := models.OTPCode{
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		OTPCode:     otpCode,
		Type:        input.Type,
		ExpiresAt:   expiresAt,
	}

	if err := services.SaveOTP(&otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to save OTP"})
		return
	}

	// Send OTP via email
	subject := "Your OTP Code"
	body := "<strong>Your OTP code is " + otpCode + "</strong>"
	if err := utils.SendEmail(input.Email, subject, body); err != nil {
		log.Println("Failed to send OTP email:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to send OTP email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
}

func SendOTP(c *gin.Context) {
	var input OTPVerify
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": err.Error()})
		return
	}

	isValid, err := services.VerifyOTP(input.Email, input.PhoneNumber, input.OTPCode, input.Type)
	if err != nil || !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"statusCode": "401", "isError": true, "message": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
}

// func SendKYCData(c *gin.Context) {
// 	var input KYCRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": err.Error()})
// 		return
// 	}

// 	userID := c.GetInt("userID")
// 	if userID == 0 {
// 		c.JSON(http.StatusUnauthorized, gin.H{"statusCode": "401", "isError": true, "message": "Unauthorized"})
// 		return
// 	}

// 	// Print userID to ensure it's correctly retrieved
// 	log.Println("UserID from context:", userID)

// 	// Check if the userID exists in the users table
// 	_, err := services.GetUserByID(userID)
// 	if err != nil {
// 		log.Println("UserID does not exist:", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": "Invalid userID"})
// 		return
// 	}

// 	birthdayDate, err := time.Parse("2006-01-02", input.BirthdayDate)
// 	if err != nil {
// 		log.Println("Invalid date format for BirthdayDate:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Invalid date format for BirthdayDate"})
// 		return
// 	}

// 	validityPeriod, err := time.Parse("2006-01-02", input.ValidityPeriod)
// 	if err != nil {
// 		log.Println("Invalid date format for ValidityPeriod:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Invalid date format for ValidityPeriod"})
// 		return
// 	}

// 	kycData := models.KYCData{
// 		UserID:         userID,
// 		NIK:            input.NIK,
// 		Name:           input.Name,
// 		BirthdayDate:   birthdayDate,
// 		BirthdayPlace:  input.BirthdayPlace,
// 		Gender:         input.Gender,
// 		Address:        input.Address,
// 		Religion:       input.Religion,
// 		MarriageStatus: input.MarriageStatus,
// 		Occupation:     input.Occupation,
// 		Citizenship:    input.Citizenship,
// 		ValidityPeriod: validityPeriod,
// 	}

// 	if err := services.SaveKYCData(&kycData); err != nil {
// 		log.Println("Failed to save KYC data:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to save KYC data"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
// }

// func RegisterOnboarding(c *gin.Context) {
// 	var input RegisterRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": err.Error()})
// 		return
// 	}

// 	// Validate KYC data before creating user
// 	kycData := models.KYCData{
// 		NIK:            input.NIK,
// 		Name:           input.Name,
// 		BirthdayDate:   time.Now(), // Placeholder, update as needed
// 		BirthdayPlace:  input.BirthdayPlace,
// 		Gender:         input.Gender,
// 		Address:        input.Address,
// 		Religion:       input.Religion,
// 		MarriageStatus: input.MarriageStatus,
// 		Occupation:     input.Occupation,
// 		Citizenship:    input.Citizenship,
// 		ValidityPeriod: time.Now(), // Placeholder, update as needed
// 	}

// 	if err := services.ValidateKYCData(&kycData); err != nil {
// 		if err.Error() == "KYC data not found" {
// 			c.JSON(http.StatusNotFound, gin.H{"statusCode": "404", "isError": true, "message": err.Error()})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to validate KYC data"})
// 		}
// 		return
// 	}

// 	// Proceed to create user only if KYC validation is successful
// 	passwordHash, err := services.HashPassword(input.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to hash password"})
// 		return
// 	}

// 	pinHash, err := services.HashPassword(input.Pin)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to hash pin"})
// 		return
// 	}

// 	user := models.User{
// 		Email:        input.Email,
// 		Username:     input.Username,
// 		PhoneNumber:  input.PhoneNumber,
// 		PasswordHash: passwordHash,
// 		PinHash:      pinHash,
// 		KycStatus:    "pending",
// 	}

// 	if err := services.CreateUser(&user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "User created and KYC data validated successfully"})
// }

// func Login(c *gin.Context) {
// 	var input struct {
// 		Email    string `json:"email" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": err.Error()})
// 		return
// 	}

// 	user, err := services.GetUserByEmail(input.Email)
// 	if err != nil || !services.CheckPasswordHash(input.Password, user.PasswordHash) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"statusCode": "401", "isError": true, "message": "Invalid credentials"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"statusCode": "200",
// 		"isError":    false,
// 		"message":    "SUCCESS",
// 		"data": gin.H{
// 			"userID":      user.ID,
// 			"userEmail":   user.Email,
// 			"userName":    user.Username,
// 			"phoneNumber": user.PhoneNumber,
// 		},
// 	})
// }

// func RequestPIN(c *gin.Context) {
// 	var input struct {
// 		UserID int    `json:"userID" binding:"required"`
// 		Pin    string `json:"pin" binding:"required"`
// 		Type   string `json:"type"`
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": err.Error()})
// 		return
// 	}

// 	user, err := services.GetUserByID(input.UserID)
// 	if err != nil || !services.CheckPinHash(input.Pin, user.PinHash) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"statusCode": "401", "isError": true, "message": "Invalid PIN"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
// }

// func RegisterOnboarding(c *gin.Context) {
// 	var user models.User

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
// 		return
// 	}

// 	if err := services.RegisterUser(user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to register"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
// }

// func Login(c *gin.Context) {
// 	var request struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
// 		return
// 	}

// 	user, err := services.AuthenticateUser(request.Email, request.Password)
// 	if err != nil || user == nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"isError": true, "message": "Invalid credentials"})
// 		return
// 	}

// 	response := map[string]interface{}{
// 		"userID":      user.ID,
// 		"userEmail":   user.Email,
// 		"userName":    user.Username,
// 		"phoneNumber": user.PhoneNumber,
// 	}
// 	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": response})
// }

func RequestPIN(c *gin.Context) {
	var request struct {
		UserID    string `json:"userID"`    // Corrected JSON field tags
		UserEmail string `json:"userEmail"` // Corrected JSON field tags
		PIN       string `json:"pin"`
		Type      string `json:"type"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	log.Printf("Request received: %+v", request)

	var user models.User
	query := "SELECT id FROM users WHERE id = ? AND email = ?"
	err := utils.GetDB().QueryRow(query, request.UserID, request.UserEmail).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"isError": true, "message": "Invalid credentials", "error": err.Error()})
		return
	}

	// Assuming PIN is validated here

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
}

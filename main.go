package main

import (
	"log"
	"os"
	"users/controllers"
	"users/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	log.Println("Loading .env file from current directory")

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}

func init() {
	log.Println("Loading .env file from current directory")

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Create a new router with default middleware (Logger and Recovery)
	r := gin.Default()

	// Initialize database
	utils.InitDB()

	// Public routes
	// r.POST("/checkEmail", controllers.CheckEmail)
	// r.POST("/requestOTP", controllers.RequestOTP)
	r.POST("/sendOTP", controllers.SendOTP)
	r.POST("/registerOnboarding", controllers.RegisterOnboarding)
	r.POST("/login", controllers.Login)
	// r.POST("/requestPIN", controllers.RequestPIN)
	// r.POST("/getBalance", controllers.GetBalance)
	// r.POST("/getBalanceHistory", controllers.GetBalanceHistory)
	// r.POST("/getCompanyGrowth", controllers.GetCompanyGrowth)
	// r.POST("/getSalesGrowth", controllers.GetSalesGrowth)
	// r.GET("/investmentReferences", controllers.GetInvestmentReferences)
	// r.POST("/cashFlowReport", controllers.GetCashFlowReport)
	// r.GET("/suppliers", controllers.GetSuppliers)
	// r.GET("/ventureCapital", controllers.GetVentureCapital)
	// r.POST("/generateQRCode", controllers.GenerateQRCode)
	// r.POST("/confirmPayment", controllers.ConfirmPayment)
	r.GET("/earthquakeData", controllers.FetchEarthquakeData)
	r.POST("/getCurrentStatus", controllers.GetCurrentStatus)
	r.POST("/getRescuersCategory", controllers.GetRescuersCategory)
	r.POST("/sendRescueDisaster", controllers.SendRescueDisaster)
	r.POST("/sendRealtimeLocation", controllers.SendRealtimeLocation)

	// Protected routes
	// authorized := r.Group("/")
	// authorized.Use(middleware.AuthMiddleware())
	// {
	// 	authorized.POST("/sendKYCdata", controllers.SendKYCData)
	// 	authorized.POST("/requestPIN", controllers.RequestPIN)
	// }

	r.Run()
}

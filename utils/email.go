package utils

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_USERNAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		os.Getenv("GMAIL_SMTP_SERVER"),
		587, // or os.Getenv("GMAIL_SMTP_PORT")
		os.Getenv("GMAIL_USERNAME"),
		os.Getenv("GMAIL_PASSWORD"),
	)
	d.SSL = false // Use STARTTLS

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}
	log.Println("Email sent successfully to", to)
	return nil
}

package services

import (
	"errors"
	"users/models"
	"users/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.User) error {
	db := utils.GetDB()

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert the user into the database
	query := "INSERT INTO users (email, name, phone_number, password) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, user.Email, user.Name, user.PhoneNumber, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(email, password string) (models.User, error) {
	db := utils.GetDB()
	var user models.User

	// Retrieve the user from the database
	query := "SELECT id, email, name, phone_number, password FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.PhoneNumber, &user.Password)
	if err != nil {
		return user, err
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}

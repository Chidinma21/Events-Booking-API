package models

import (
	"errors"

	"github.com/Chidinma21/Events-Booking-API/db"
	"github.com/Chidinma21/Events-Booking-API/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT password FROM users where email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}
	
	valid := utils.ComparePassword(u.Password, retrievedPassword)
	if !valid { 
		return errors.New("credentials invalid")
	}

	return nil
}

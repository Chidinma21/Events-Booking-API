package models

import (
	"errors"

	"github.com/Chidinma21/Events-Booking-API/db"
)

type Registration struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id" binding:"required"`
	EventID int64 `json:"event_id" binding:"required"`
} 

func (r Registration) Save() error {
	query := `INSERT INTO registrations(user_id, event_id) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.UserID, r.EventID)

	if err != nil {
		return err
	}
	return nil
}

func (r Registration) CancelRegistration() error {
	query := `DELETE FROM registrations WHERE user_id = ? AND event_id = ?` 
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(r.UserID, r.EventID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no registration found to delete")
	}

	return nil
}
package handlers

import (
	"database/sql"
	"errors"

	"forum/security"
	"forum/validate"
)

func CreateAccount(db *sql.DB, email string, username string, password string) (bool, error) {

	if !validate.IsValidUsername(username) {
		return false, errors.New("invalid username")
	}

	if !validate.IsUniqueUsername(db, username) {
		return false, errors.New("username already exists")
	}

	if !validate.IsValidEmail(email) {
		return false, errors.New("invalid email")
	}

	if !validate.IsUniqueEmail(db, email) {
		return false, errors.New("email already exists")
	}

	if !validate.IsValidPassword(password) {
		return false, errors.New("invalid password")
	}

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return false, err
	}

	query := `
		INSERT INTO users (email, username, password)
		VALUES (?, ?, ?)
	`

	_, err = db.Exec(query, email, username, hashedPassword)
	if err != nil {
		return false, err
	}

	return true, nil
}
package handlers

import (
	"database/sql"
	"errors"

	"forum/auth"
)

func LogAccount(db *sql.DB, email string, password string) (bool, error) {

	user, err := auth.GetUserByEmail(db, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("utilisateur introuvable")
		}
		return false, err
	}

	if user.Password != password {
		return false, errors.New("mot de passe incorrect")
	}

	return true, nil
}
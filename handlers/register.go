package users

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
)

func CreateAccount(db *sql.DB, email string, username string, password string) (bool, error) {

	if strings.TrimSpace(email) == "" {
		return false, errors.New("email obligatoire")
	}

	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return false, errors.New("email invalide")
	}

	if strings.TrimSpace(username) == "" {
		return false, errors.New("nom d'utilisateur obligatoire")
	}

	if len(username) < 3 || len(username) > 30 {
		return false, errors.New("le nom d'utilisateur doit contenir entre 3 et 30 caractères")
	}

	usernameRegex := `^[a-zA-Z0-9_]+$`
	if !regexp.MustCompile(usernameRegex).MatchString(username) {
		return false, errors.New("nom d'utilisateur invalide (lettres, chiffres et _ uniquement)")
	}

	if len(password) < 8 {
		return false, errors.New("le mot de passe doit contenir au moins 8 caractères")
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber {
		return false, errors.New("le mot de passe doit contenir une majuscule, une minuscule et un chiffre")
	}

	query := `
		INSERT INTO users (email, username, password)
		VALUES (?, ?, ?)
	`

	_, err := db.Exec(query, email, username, password)
	if err != nil {
		return false, err
	}

	return true, nil
}
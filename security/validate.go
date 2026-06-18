package security

import "strings"

func IsValidUsername(username string) bool {
	var specialChars = "-_"
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	for i := 0; i < len(username); i++ {
		if (username[i] >= 'A' && username[i] <= 'Z') || (username[i] >= 'a' && username[i] <= 'z') || (username[i] >= '0' && username[i] <= '9') || strings.ContainsRune(specialChars, rune(username[i])) {
			continue
		} else {
			return false
		}
	}
	return true
}

func IsUniqueUsername(base []string, newname string) bool {
	for _, x := range base {
		if x == newname {
			return false
		}
	}
	return true
}

func IsValidPassword(password string) bool {
	var specialChars = "!()-.?[]_`~;:@#$%^&*+="
	if len(password) < 8 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	for i := 0; i < len(password); i++ {
		if password[i] >= 'A' && password[i] <= 'Z' {
			hasUpper = true
		} else if password[i] >= 'a' && password[i] <= 'z' {
			hasLower = true
		} else if password[i] >= '0' && password[i] <= '9' {
			hasDigit = true
		} else if strings.ContainsRune(specialChars, rune(password[i])) {
			hasSpecial = true
		}
	}
	if hasUpper && hasLower && hasDigit && hasSpecial == true {
		return true
	}
	return false
}

func IsValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			return true
		}
	}
	return false
}

func IsUniqueEmail(base []string, newmail string) bool {
	for _, x := range base {
		if x == newmail {
			return false
		}
	}
	return true
}
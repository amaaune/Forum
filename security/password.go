package security

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(hash), err
}

func CheckPassword(testpassword, truepassword string) bool {
    hash, _ := HashPassword(truepassword)
   
	testPassword, _ := HashPassword(testpassword)

	if testPassword == hash {
		fmt.Println("Mot de passe correct")
		return true
	} else {
		fmt.Println("Mot de passe incorrect")
		return false
	}
}
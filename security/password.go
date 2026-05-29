package security

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(hash), err
}

func checkPassword(testpassword, truepassword string) bool {
    hash, _ := hashPassword(truepassword)
   
	testPassword, _ := hashPassword(testpassword)

	if testPassword == hash {
		fmt.Println("Mot de passe correct")
		return true
	} else {
		fmt.Println("Mot de passe incorrect")
		return false
	}
}
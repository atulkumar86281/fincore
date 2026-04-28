package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Returns the bcrypt hash of the password
func HashedPassword(pass string) (string,error){
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
	if err != nil{
		return "",fmt.Errorf("failed to hash password %w",err)
	}

	return string(hashedPass),nil
}

func CheckPassword(pass string, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass),[]byte(pass))
}
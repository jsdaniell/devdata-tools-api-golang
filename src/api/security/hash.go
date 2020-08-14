package security

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(pass string) ([]byte, error){

	var crypt, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return make([]byte, 0), err
	}
	return crypt, nil
}

func VerifyPass(hashedPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}
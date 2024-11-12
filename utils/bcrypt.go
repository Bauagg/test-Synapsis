package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) *string {
	result, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err.Error())
	}

	hashedPassword := string(result)
	return &hashedPassword
}

func VerifikasiHashPassword(password, dbPassword *string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*dbPassword), []byte(*password))

	return err
}
